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

// loadFixture loads the testdata fixture package into a decorated tree with full
// type information, the same way the engine loads the real provider packages.
func loadFixture(t *testing.T) *decorator.Package {
	t.Helper()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedDeps | packages.NeedTypes |
			packages.NeedTypesInfo | packages.NeedTypesSizes | packages.NeedSyntax,
		Dir:  ".",
		Fset: token.NewFileSet(),
	}
	pkgs, err := decorator.Load(cfg, "./testdata/fixture")
	require.NoError(t, err)
	require.Len(t, pkgs, 1)
	require.Empty(t, pkgs[0].Errors, "fixture must type-check cleanly")
	return pkgs[0]
}

// detectAll runs every detection stage over a single-package fixture, mirroring
// Pass.Detect but without the engine's hardcoded multi-package layout.
func detectAll(p *decorator.Package) []engine.Target {
	var targets []engine.Target
	marked := map[fieldRef]bool{}
	sdkRefs := map[fieldRef]bool{}

	detectModel(p, &targets, marked)
	detectSchema(p, &targets)
	_ = detectRefresh(p, &targets, sdkRefs)
	detectBuild(p, &targets, marked)
	detectSDK(p, &targets, sdkRefs)
	return targets
}

func TestDetect(t *testing.T) {
	p := loadFixture(t)
	targets := detectAll(p)

	counts := map[kind]int{}
	for _, tg := range targets {
		counts[tg.Data.(payload).kind]++
	}

	require.Equal(t, 1, counts[kindModel], "model field")
	require.Equal(t, 1, counts[kindSchema], "schema attribute")
	require.Equal(t, 1, counts[kindRefresh], "refresh site")
	require.Equal(t, 1, counts[kindBuild], "build site")
	require.Equal(t, 1, counts[kindSDKField], "sdk field")
	require.Equal(t, 1, counts[kindSDKGetter], "sdk getter")
}

func TestApplyGolden(t *testing.T) {
	p := loadFixture(t)
	targets := detectAll(p)
	require.NotEmpty(t, targets)

	pass := New()
	for _, tg := range targets {
		require.NoError(t, pass.apply(tg), "apply %s", tg.Name)
	}

	var file *dst.File
	for _, f := range p.Syntax {
		file = f
		break
	}
	require.NotNil(t, file)

	got, err := engine.RenderFile(p, file)
	require.NoError(t, err)

	golden := filepath.Join("testdata", "fixture", "expected.go.golden")
	if *update {
		require.NoError(t, os.WriteFile(golden, got, 0o644))
		return
	}
	want, err := os.ReadFile(golden)
	require.NoError(t, err, "run with -update to create the golden file")
	require.Equal(t, string(want), string(got))
}
