package lucide

import x "github.com/bloxui/blox"

// SquareUser creates a Square User Lucide icon.
func SquareUser(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-user", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("3"))),
		x.Child(x.Path(x.D("M7 21v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2"))),
	)
	return x.Svg(svgArgs...)
}
