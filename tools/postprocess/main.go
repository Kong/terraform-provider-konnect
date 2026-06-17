// Command postprocess re-applies the deterministic transformations the
// generated Terraform provider needs but Speakeasy cannot emit.
//
// It runs immediately after `speakeasy run` (see `make speakeasy`): it loads the
// generated provider packages, runs each registered Pass, formats the result
// with goimports and gates on `go build ./...`. The transformations are
// idempotent, so a second run over already-processed code is a no-op — which is
// what the CI clean-diff check relies on.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kong/terraform-provider-konnect/v3/tools/postprocess/engine"
	"github.com/kong/terraform-provider-konnect/v3/tools/postprocess/passes/dynamictype"
)

func main() {
	dir := flag.String("dir", ".", "module root to post-process")
	flag.Parse()

	reg := &engine.Registry{}
	reg.Register(dynamictype.New())

	if err := engine.Run(*dir, reg); err != nil {
		fmt.Fprintf(os.Stderr, "postprocess: %v\n", err)
		os.Exit(1)
	}
}
