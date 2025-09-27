package lucide

import (
	html "github.com/plainkit/html"
)

// ChartGantt creates a Chart Gantt Lucide icon.
func ChartGantt(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-gantt", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 6h8"))),
		html.Child(html.SvgPath(html.AD("M12 16h6"))),
		html.Child(html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16"))),
		html.Child(html.SvgPath(html.AD("M8 11h7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
