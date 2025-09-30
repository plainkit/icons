package lucide

import (
	html "github.com/plainkit/html"
)

// ChartNoAxesGantt creates a Chart No Axes Gantt Lucide icon.
func ChartNoAxesGantt(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-no-axes-gantt", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 5h12")),
		html.SvgPath(html.AD("M4 12h10")),
		html.SvgPath(html.AD("M12 19h8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
