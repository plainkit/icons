package lucide

import x "github.com/bloxui/blox"

// Share2 creates a Share 2 Lucide icon.
func Share2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-share-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("18"), x.Cy("5"), x.R("3"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("12"), x.R("3"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("19"), x.R("3"))),
		x.Child(x.Line(x.X1("8.59"), x.X2("15.42"), x.Y1("13.51"), x.Y2("17.49"))),
		x.Child(x.Line(x.X1("15.41"), x.X2("8.59"), x.Y1("6.51"), x.Y2("10.49"))),
	)
	return x.Svg(svgArgs...)
}
