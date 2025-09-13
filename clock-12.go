package lucide

import x "github.com/bloxui/blox"

// Clock12 creates a Clock 12 Lucide icon.
func Clock12(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-clock-12", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
