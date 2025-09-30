package lucide

import (
	html "github.com/plainkit/html"
)

// SquareChartGantt creates a Square Chart Gantt Lucide icon.
func SquareChartGantt(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-chart-gantt", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M9 8h7")),
		html.SvgPath(html.AD("M8 12h6")),
		html.SvgPath(html.AD("M11 16h5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
