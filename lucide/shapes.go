package lucide

import x "github.com/plainkit/html"

// Shapes creates a Shapes Lucide icon.
func Shapes(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shapes", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8.3 10a.7.7 0 0 1-.626-1.079L11.4 3a.7.7 0 0 1 1.198-.043L16.3 8.9a.7.7 0 0 1-.572 1.1Z"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("3"), x.Y("14"), x.Rx("1"))),
		x.Child(x.Circle(x.Cx("17.5"), x.Cy("17.5"), x.R("3.5"))),
	)
	return x.Svg(svgArgs...)
}
