package lucide

import x "github.com/bloxui/blox"

// Repeat2 creates a Repeat 2 Lucide icon.
func Repeat2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-repeat-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2 9 3-3 3 3"))),
		x.Child(x.Path(x.D("M13 18H7a2 2 0 0 1-2-2V6"))),
		x.Child(x.Path(x.D("m22 15-3 3-3-3"))),
		x.Child(x.Path(x.D("M11 6h6a2 2 0 0 1 2 2v10"))),
	)
	return x.Svg(svgArgs...)
}
