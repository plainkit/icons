package lucide

import x "github.com/plainkit/html"

// Voicemail creates a Voicemail Lucide icon.
func Voicemail(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-voicemail", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("6"), x.Cy("12"), x.R("4"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("12"), x.R("4"))),
		x.Child(x.Line(x.X1("6"), x.X2("18"), x.Y1("16"), x.Y2("16"))),
	)
	return x.Svg(svgArgs...)
}
