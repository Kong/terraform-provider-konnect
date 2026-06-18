package dynamictype

import (
	"flag"
	"go/token"
	"os"
	"path/filepath"
	"testing"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/kong/terraform-provider-konnect/v3/tools/postprocess/engine"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"
)

var update = flag.Bool("update", false, "update golden files")

const (
	fixtureProvider = "github.com/kong/terraform-provider-konnect/v3/tools/postprocess/passes/dynamictype/testdata/fixture/provider"
	fixtureShared   = "github.com/kong/terraform-provider-konnect/v3/tools/postprocess/passes/dynamictype/testdata/fixture/shared"
)

// loadFixture loads the two fixture packages (provider + shared) into decorated
// trees with full type information, the same way the engine loads the real
// provider packages. The package separation matters: detectSDK must run only
// over the shared package, exactly as in the real layout.
func loadFixture(t *testing.T) (providerPkg, sharedPkg *decorator.Package) {
	t.Helper()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedDeps | packages.NeedTypes |
			packages.NeedTypesInfo | packages.NeedTypesSizes | packages.NeedSyntax,
		Dir:  ".",
		Fset: token.NewFileSet(),
	}
	pkgs, err := decorator.Load(cfg, "./testdata/fixture/provider", "./testdata/fixture/shared")
	require.NoError(t, err)
	for _, p := range pkgs {
		require.Empty(t, p.Errors, "fixture %s must type-check cleanly", p.PkgPath)
		switch p.PkgPath {
		case fixtureProvider:
			providerPkg = p
		case fixtureShared:
			sharedPkg = p
		}
	}
	require.NotNil(t, providerPkg)
	require.NotNil(t, sharedPkg)
	return providerPkg, sharedPkg
}

// detectAll runs every detection stage across the two fixture packages, mirroring
// Pass.Detect but with the fixture's package paths instead of the engine's
// hardcoded layout. Provider-side stages run over providerPkg; detectSDK runs
// over sharedPkg only.
func detectAll(providerPkg, sharedPkg *decorator.Package) []engine.Target {
	var targets []engine.Target
	marked := map[fieldRef]bool{}
	sdkRefs := map[fieldRef]bool{}
	buildNames := map[[2]string]bool{}

	detectModel(providerPkg, &targets, marked)
	detectSchema(providerPkg, &targets)
	_ = detectRefresh(providerPkg, &targets, sdkRefs)
	detectBuild(providerPkg, &targets, marked, buildNames)
	detectSDK(sharedPkg, &targets, sdkRefs, buildNames)
	return targets
}

func TestDetect(t *testing.T) {
	providerPkg, sharedPkg := loadFixture(t)
	targets := detectAll(providerPkg, sharedPkg)

	counts := map[kind]int{}
	for _, tg := range targets {
		counts[tg.Data.(payload).kind]++
	}

	// Two marked fields: Redis.Port (pointer, refresh-linked) and
	// ACLRule.ResourceNames (value, build-name-linked).
	require.Equal(t, 2, counts[kindModel], "model fields")
	require.Equal(t, 2, counts[kindSchema], "schema attributes")
	require.Equal(t, 1, counts[kindRefresh], "refresh site (only Redis.Port has one)")
	require.Equal(t, 2, counts[kindBuild], "build sites")
	require.Equal(t, 2, counts[kindSDKField], "sdk fields")
	require.Equal(t, 2, counts[kindSDKGetter], "sdk getters")
}

func TestApplyGolden(t *testing.T) {
	providerPkg, sharedPkg := loadFixture(t)
	targets := detectAll(providerPkg, sharedPkg)
	require.NotEmpty(t, targets)

	pass := New()
	for _, tg := range targets {
		require.NoError(t, pass.apply(tg), "apply %s", tg.Name)
	}

	for _, tc := range []struct {
		pkg    *decorator.Package
		golden string
	}{
		{providerPkg, "provider/expected.go.golden"},
		{sharedPkg, "shared/expected.go.golden"},
	} {
		var file *dst.File
		for _, f := range tc.pkg.Syntax {
			file = f
			break
		}
		require.NotNil(t, file)

		got, err := engine.RenderFile(tc.pkg, file)
		require.NoError(t, err)

		golden := filepath.Join("testdata", "fixture", tc.golden)
		if *update {
			require.NoError(t, os.WriteFile(golden, got, 0o644))
			continue
		}
		want, err := os.ReadFile(golden)
		require.NoError(t, err, "run with -update to create %s", tc.golden)
		require.Equal(t, string(want), string(got), tc.golden)
	}
}
