package lucide

import x "github.com/plainkit/html"

// Disc3 creates a Disc 3 Lucide icon.
func Disc3(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-disc-3", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M6 12c0-1.7.7-3.2 1.8-4.2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
		x.Child(x.Path(x.D("M18 12c0 1.7-.7 3.2-1.8 4.2"))),
	)
	return x.Svg(svgArgs...)
}
