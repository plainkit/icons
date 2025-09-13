package lucide

import x "github.com/bloxui/blox"

// Clock10 creates a Clock 10 Lucide icon.
func Clock10(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-clock-10", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6l-4-2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
