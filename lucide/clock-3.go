package lucide

import x "github.com/bloxui/blox"

// Clock3 creates a Clock 3 Lucide icon.
func Clock3(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-clock-3", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6h4"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
