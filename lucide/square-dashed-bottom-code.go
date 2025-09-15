package lucide

import x "github.com/plainkit/blox"

// SquareDashedBottomCode creates a Square Dashed Bottom Code Lucide icon.
func SquareDashedBottomCode(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-dashed-bottom-code", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 9.5 8 12l2 2.5"))),
		x.Child(x.Path(x.D("M14 21h1"))),
		x.Child(x.Path(x.D("m14 9.5 2 2.5-2 2.5"))),
		x.Child(x.Path(x.D("M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M9 21h1"))),
	)
	return x.Svg(svgArgs...)
}
