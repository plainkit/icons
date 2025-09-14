package lucide

import x "github.com/bloxui/blox"

// CircleArrowUp creates a Circle Arrow Up Lucide icon.
func CircleArrowUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-arrow-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m16 12-4-4-4 4"))),
		x.Child(x.Path(x.D("M12 16V8"))),
	)
	return x.Svg(svgArgs...)
}
