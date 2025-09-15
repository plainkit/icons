package lucide

import x "github.com/plainkit/html"

// CirclePower creates a Circle Power Lucide icon.
func CirclePower(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-power", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 7v4"))),
		x.Child(x.Path(x.D("M7.998 9.003a5 5 0 1 0 8-.005"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
