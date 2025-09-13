package lucide

import x "github.com/bloxui/blox"

// SquareSlash creates a Square Slash Lucide icon.
func SquareSlash(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-square-slash", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Line(x.X1("9"), x.X2("15"), x.Y1("15"), x.Y2("9"))),
	)
	return x.Svg(svgArgs...)
}
