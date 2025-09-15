package lucide

import x "github.com/plainkit/html"

// Scale3d creates a Scale 3d Lucide icon.
func Scale3d(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-scale-3d", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 7v11a1 1 0 0 0 1 1h11"))),
		x.Child(x.Path(x.D("M5.293 18.707 11 13"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("19"), x.R("2"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("5"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
