package lucide

import x "github.com/bloxui/blox"

// Undo creates a Undo Lucide icon.
func Undo(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-undo", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 7v6h6"))),
		x.Child(x.Path(x.D("M21 17a9 9 0 0 0-9-9 9 9 0 0 0-6 2.3L3 13"))),
	)
	return x.Svg(svgArgs...)
}
