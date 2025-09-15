package lucide

import x "github.com/plainkit/blox"

// Clock1 creates a Clock 1 Lucide icon.
func Clock1(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock-1", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6l2-4"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
