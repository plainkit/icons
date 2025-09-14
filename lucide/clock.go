package lucide

import x "github.com/bloxui/blox"

// Clock creates a Clock Lucide icon.
func Clock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6l4 2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
