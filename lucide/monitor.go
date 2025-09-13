package lucide

import x "github.com/bloxui/blox"

// Monitor creates a Monitor Lucide icon.
func Monitor(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-monitor", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Line(x.X1("8"), x.X2("16"), x.Y1("21"), x.Y2("21"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("17"), x.Y2("21"))),
	)
	return x.Svg(svgArgs...)
}
