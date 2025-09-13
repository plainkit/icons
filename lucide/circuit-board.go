package lucide

import x "github.com/bloxui/blox"

// CircuitBoard creates a Circuit Board Lucide icon.
func CircuitBoard(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circuit-board", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M11 9h4a2 2 0 0 0 2-2V3"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("9"), x.R("2"))),
		x.Child(x.Path(x.D("M7 21v-4a2 2 0 0 1 2-2h4"))),
		x.Child(x.Circle(x.Cx("15"), x.Cy("15"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
