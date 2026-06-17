package engine

import (
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

// Target is one unit of work produced by a Pass's Detect and consumed by its
// Apply. The engine treats Data opaquely; each pass defines and type-asserts its
// own payload. Pkg/File let the engine (and the pass) know which decorated file
// to mutate and mark as touched.
type Target struct {
	// Name is a human-readable label for the unit of work (used in logs/errors).
	Name string
	// Pkg is the package whose dst tree the target lives in.
	Pkg *decorator.Package
	// File is the decorated file the target lives in.
	File *dst.File
	// Data is the pass-specific payload (located nodes, metadata).
	Data any
}
