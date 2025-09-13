package lucide

import x "github.com/bloxui/blox"

// Codesandbox creates a Codesandbox Lucide icon.
func Codesandbox(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-codesandbox", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"))),
		x.Child(x.Polyline(x.Points("7.5 4.21 12 6.81 16.5 4.21"))),
		x.Child(x.Polyline(x.Points("7.5 19.79 7.5 14.6 3 12"))),
		x.Child(x.Polyline(x.Points("21 12 16.5 14.6 16.5 19.79"))),
		x.Child(x.Polyline(x.Points("3.27 6.96 12 12.01 20.73 6.96"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("22.08"), x.Y2("12"))),
	)
	return x.Svg(svgArgs...)
}
