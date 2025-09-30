package lucide

import (
	html "github.com/plainkit/html"
)

// ChartNoAxesColumn creates a Chart No Axes Column Lucide icon.
func ChartNoAxesColumn(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-no-axes-column", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 21v-6")),
		html.SvgPath(html.AD("M12 21V3")),
		html.SvgPath(html.AD("M19 21V9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
