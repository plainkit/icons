package lucide

import (
	html "github.com/plainkit/html"
)

// ChartColumnIncreasing creates a Chart Column Increasing Lucide icon.
func ChartColumnIncreasing(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-column-increasing", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13 17V9"))),
		html.Child(html.SvgPath(html.AD("M18 17V5"))),
		html.Child(html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16"))),
		html.Child(html.SvgPath(html.AD("M8 17v-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
