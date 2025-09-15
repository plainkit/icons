package lucide

import x "github.com/plainkit/blox"

// Diameter creates a Diameter Lucide icon.
func Diameter(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-diameter", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("19"), x.Cy("19"), x.R("2"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("5"), x.R("2"))),
		x.Child(x.Path(x.D("M6.48 3.66a10 10 0 0 1 13.86 13.86"))),
		x.Child(x.Path(x.D("m6.41 6.41 11.18 11.18"))),
		x.Child(x.Path(x.D("M3.66 6.48a10 10 0 0 0 13.86 13.86"))),
	)
	return x.Svg(svgArgs...)
}
