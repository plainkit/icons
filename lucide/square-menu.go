package lucide

import x "github.com/plainkit/blox"

// SquareMenu creates a Square Menu Lucide icon.
func SquareMenu(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-menu", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M7 8h10"))),
		x.Child(x.Path(x.D("M7 12h10"))),
		x.Child(x.Path(x.D("M7 16h10"))),
	)
	return x.Svg(svgArgs...)
}
