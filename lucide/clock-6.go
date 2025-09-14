package lucide

import x "github.com/bloxui/blox"

// Clock6 creates a Clock 6 Lucide icon.
func Clock6(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock-6", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v10"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
