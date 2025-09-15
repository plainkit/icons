package lucide

import x "github.com/plainkit/blox"

// SquareAsterisk creates a Square Asterisk Lucide icon.
func SquareAsterisk(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-asterisk", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 8v8"))),
		x.Child(x.Path(x.D("m8.5 14 7-4"))),
		x.Child(x.Path(x.D("m8.5 10 7 4"))),
	)
	return x.Svg(svgArgs...)
}
