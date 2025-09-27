package lucide

import (
	html "github.com/plainkit/html"
)

// ChartScatter creates a Chart Scatter Lucide icon.
func ChartScatter(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-scatter", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("7.5"), html.ACy("7.5"), html.AR(".5"), html.AFill("currentColor"))),
		html.Child(html.SvgCircle(html.ACx("18.5"), html.ACy("5.5"), html.AR(".5"), html.AFill("currentColor"))),
		html.Child(html.SvgCircle(html.ACx("11.5"), html.ACy("11.5"), html.AR(".5"), html.AFill("currentColor"))),
		html.Child(html.SvgCircle(html.ACx("7.5"), html.ACy("16.5"), html.AR(".5"), html.AFill("currentColor"))),
		html.Child(html.SvgCircle(html.ACx("17.5"), html.ACy("14.5"), html.AR(".5"), html.AFill("currentColor"))),
		html.Child(html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
