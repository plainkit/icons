package lucide

import (
	html "github.com/plainkit/html"
)

// ChartBarDecreasing creates a Chart Bar Decreasing Lucide icon.
func ChartBarDecreasing(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-bar-decreasing", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16")),
		html.SvgPath(html.AD("M7 11h8")),
		html.SvgPath(html.AD("M7 16h3")),
		html.SvgPath(html.AD("M7 6h12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
