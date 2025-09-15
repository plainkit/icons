package lucide

import x "github.com/plainkit/html"

// Bubbles creates a Bubbles Lucide icon.
func Bubbles(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bubbles", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7.2 14.8a2 2 0 0 1 2 2"))),
		x.Child(x.Circle(x.Cx("18.5"), x.Cy("8.5"), x.R("3.5"))),
		x.Child(x.Circle(x.Cx("7.5"), x.Cy("16.5"), x.R("5.5"))),
		x.Child(x.Circle(x.Cx("7.5"), x.Cy("4.5"), x.R("2.5"))),
	)
	return x.Svg(svgArgs...)
}
