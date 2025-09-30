package lucide

import (
	html "github.com/plainkit/html"
)

// ChartNoAxesColumnIncreasing creates a Chart No Axes Column Increasing Lucide icon.
func ChartNoAxesColumnIncreasing(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-no-axes-column-increasing", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 21v-6")),
		html.SvgPath(html.AD("M12 21V9")),
		html.SvgPath(html.AD("M19 21V3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
