package lucide

import x "github.com/plainkit/blox"

// SquareLibrary creates a Square Library Lucide icon.
func SquareLibrary(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-library", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M7 7v10"))),
		x.Child(x.Path(x.D("M11 7v10"))),
		x.Child(x.Path(x.D("m15 7 2 10"))),
	)
	return x.Svg(svgArgs...)
}
