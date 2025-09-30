package lucide

import (
	html "github.com/plainkit/html"
)

// ChartBarIncreasing creates a Chart Bar Increasing Lucide icon.
func ChartBarIncreasing(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-bar-increasing", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16")),
		html.SvgPath(html.AD("M7 11h8")),
		html.SvgPath(html.AD("M7 16h12")),
		html.SvgPath(html.AD("M7 6h3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
