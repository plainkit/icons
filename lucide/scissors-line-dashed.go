package lucide

import x "github.com/bloxui/blox"

// ScissorsLineDashed creates a Scissors Line Dashed Lucide icon.
func ScissorsLineDashed(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-scissors-line-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5.42 9.42 8 12"))),
		x.Child(x.Circle(x.Cx("4"), x.Cy("8"), x.R("2"))),
		x.Child(x.Path(x.D("m14 6-8.58 8.58"))),
		x.Child(x.Circle(x.Cx("4"), x.Cy("16"), x.R("2"))),
		x.Child(x.Path(x.D("M10.8 14.8 14 18"))),
		x.Child(x.Path(x.D("M16 12h-2"))),
		x.Child(x.Path(x.D("M22 12h-2"))),
	)
	return x.Svg(svgArgs...)
}
