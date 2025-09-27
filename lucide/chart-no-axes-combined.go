package lucide

import (
	html "github.com/plainkit/html"
)

// ChartNoAxesCombined creates a Chart No Axes Combined Lucide icon.
func ChartNoAxesCombined(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-no-axes-combined", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 16v5"))),
		html.Child(html.SvgPath(html.AD("M16 14v7"))),
		html.Child(html.SvgPath(html.AD("M20 10v11"))),
		html.Child(html.SvgPath(html.AD("m22 3-8.646 8.646a.5.5 0 0 1-.708 0L9.354 8.354a.5.5 0 0 0-.707 0L2 15"))),
		html.Child(html.SvgPath(html.AD("M4 18v3"))),
		html.Child(html.SvgPath(html.AD("M8 14v7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
