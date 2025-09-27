package lucide

import (
	html "github.com/plainkit/html"
)

// ChartNoAxesColumnDecreasing creates a Chart No Axes Column Decreasing Lucide icon.
func ChartNoAxesColumnDecreasing(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-no-axes-column-decreasing", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 21V3"))),
		html.Child(html.SvgPath(html.AD("M12 21V9"))),
		html.Child(html.SvgPath(html.AD("M19 21v-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
