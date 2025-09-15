package lucide

import x "github.com/plainkit/blox"

// Shrimp creates a Shrimp Lucide icon.
func Shrimp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shrimp", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 12h.01"))),
		x.Child(x.Path(x.D("M13 22c.5-.5 1.12-1 2.5-1-1.38 0-2-.5-2.5-1"))),
		x.Child(x.Path(x.D("M14 2a3.28 3.28 0 0 1-3.227 1.798l-6.17-.561A2.387 2.387 0 1 0 4.387 8H15.5a1 1 0 0 1 0 13 1 1 0 0 0 0-5H12a7 7 0 0 1-7-7V8"))),
		x.Child(x.Path(x.D("M14 8a8.5 8.5 0 0 1 0 8"))),
		x.Child(x.Path(x.D("M16 16c2 0 4.5-4 4-6"))),
	)
	return x.Svg(svgArgs...)
}
