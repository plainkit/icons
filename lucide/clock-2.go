package lucide

import x "github.com/bloxui/blox"

// Clock2 creates a Clock 2 Lucide icon.
func Clock2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6l4-2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
