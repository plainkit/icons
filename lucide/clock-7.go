package lucide

import x "github.com/plainkit/blox"

// Clock7 creates a Clock 7 Lucide icon.
func Clock7(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock-7", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6l-2 4"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
