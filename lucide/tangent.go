package lucide

import x "github.com/plainkit/html"

// Tangent creates a Tangent Lucide icon.
func Tangent(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tangent", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("17"), x.Cy("4"), x.R("2"))),
		x.Child(x.Path(x.D("M15.59 5.41 5.41 15.59"))),
		x.Child(x.Circle(x.Cx("4"), x.Cy("17"), x.R("2"))),
		x.Child(x.Path(x.D("M12 22s-4-9-1.5-11.5S22 12 22 12"))),
	)
	return x.Svg(svgArgs...)
}
