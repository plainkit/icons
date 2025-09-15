package lucide

import x "github.com/plainkit/blox"

// CircleEllipsis creates a Circle Ellipsis Lucide icon.
func CircleEllipsis(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-ellipsis", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M17 12h.01"))),
		x.Child(x.Path(x.D("M12 12h.01"))),
		x.Child(x.Path(x.D("M7 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
