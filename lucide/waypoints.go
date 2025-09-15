package lucide

import x "github.com/plainkit/blox"

// Waypoints creates a Waypoints Lucide icon.
func Waypoints(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-waypoints", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("4.5"), x.R("2.5"))),
		x.Child(x.Path(x.D("m10.2 6.3-3.9 3.9"))),
		x.Child(x.Circle(x.Cx("4.5"), x.Cy("12"), x.R("2.5"))),
		x.Child(x.Path(x.D("M7 12h10"))),
		x.Child(x.Circle(x.Cx("19.5"), x.Cy("12"), x.R("2.5"))),
		x.Child(x.Path(x.D("m13.8 17.7 3.9-3.9"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("19.5"), x.R("2.5"))),
	)
	return x.Svg(svgArgs...)
}
