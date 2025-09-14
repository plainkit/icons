package lucide

import x "github.com/bloxui/blox"

// Clock8 creates a Clock 8 Lucide icon.
func Clock8(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock-8", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6l-4 2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
