package lucide

import x "github.com/bloxui/blox"

// Navigation2Off creates a Navigation 2 Off Lucide icon.
func Navigation2Off(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-navigation-2-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9.31 9.31 5 21l7-4 7 4-1.17-3.17"))),
		x.Child(x.Path(x.D("M14.53 8.88 12 2l-1.17 3.17"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("2"), x.Y2("22"))),
	)
	return x.Svg(svgArgs...)
}
