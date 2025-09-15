package lucide

import x "github.com/plainkit/html"

// Chromium creates a Chromium Lucide icon.
func Chromium(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chromium", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.88 21.94 15.46 14"))),
		x.Child(x.Path(x.D("M21.17 8H12"))),
		x.Child(x.Path(x.D("M3.95 6.06 8.54 14"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("4"))),
	)
	return x.Svg(svgArgs...)
}
