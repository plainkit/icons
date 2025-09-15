package lucide

import x "github.com/plainkit/blox"

// SquarePilcrow creates a Square Pilcrow Lucide icon.
func SquarePilcrow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-pilcrow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 12H9.5a2.5 2.5 0 0 1 0-5H17"))),
		x.Child(x.Path(x.D("M12 7v10"))),
		x.Child(x.Path(x.D("M16 7v10"))),
	)
	return x.Svg(svgArgs...)
}
