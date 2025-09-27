package lucide

import (
	html "github.com/plainkit/html"
)

// ChartColumnDecreasing creates a Chart Column Decreasing Lucide icon.
func ChartColumnDecreasing(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-column-decreasing", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13 17V9"))),
		html.Child(html.SvgPath(html.AD("M18 17v-3"))),
		html.Child(html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16"))),
		html.Child(html.SvgPath(html.AD("M8 17V5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
