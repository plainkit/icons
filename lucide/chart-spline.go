package lucide

import (
	html "github.com/plainkit/html"
)

// ChartSpline creates a Chart Spline Lucide icon.
func ChartSpline(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-spline", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16")),
		html.SvgPath(html.AD("M7 16c.5-2 1.5-7 4-7 2 0 2 3 4 3 2.5 0 4.5-5 5-7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
