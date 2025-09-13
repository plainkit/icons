package lucide

import x "github.com/bloxui/blox"

// Gamepad creates a Gamepad Lucide icon.
func Gamepad(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-gamepad", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("6"), x.X2("10"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("8"), x.X2("8"), x.Y1("10"), x.Y2("14"))),
		x.Child(x.Line(x.X1("15"), x.X2("15.01"), x.Y1("13"), x.Y2("13"))),
		x.Child(x.Line(x.X1("18"), x.X2("18.01"), x.Y1("11"), x.Y2("11"))),
		x.Child(x.Rect(x.X("2"), x.Y("6"), x.RectWidth("20"), x.RectHeight("12"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
