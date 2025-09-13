package lucide

import x "github.com/bloxui/blox"

// Clock9 creates a Clock 9 Lucide icon.
func Clock9(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-clock-9", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6H8"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
