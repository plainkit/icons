package lucide

import x "github.com/plainkit/html"

// Tablets creates a Tablets Lucide icon.
func Tablets(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tablets", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("7"), x.Cy("7"), x.R("5"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("17"), x.R("5"))),
		x.Child(x.Path(x.D("M12 17h10"))),
		x.Child(x.Path(x.D("m3.46 10.54 7.08-7.08"))),
	)
	return x.Svg(svgArgs...)
}
