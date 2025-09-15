package lucide

import x "github.com/plainkit/html"

// ArrowUpFromDot creates a Arrow Up From Dot Lucide icon.
func ArrowUpFromDot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-from-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m5 9 7-7 7 7"))),
		x.Child(x.Path(x.D("M12 16V2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("21"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
