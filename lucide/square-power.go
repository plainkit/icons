package lucide

import x "github.com/plainkit/blox"

// SquarePower creates a Square Power Lucide icon.
func SquarePower(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-power", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 7v4"))),
		x.Child(x.Path(x.D("M7.998 9.003a5 5 0 1 0 8-.005"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
