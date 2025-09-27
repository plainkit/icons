package lucide

import (
	html "github.com/plainkit/html"
)

// ChartBar creates a Chart Bar Lucide icon.
func ChartBar(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-bar", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16"))),
		html.Child(html.SvgPath(html.AD("M7 16h8"))),
		html.Child(html.SvgPath(html.AD("M7 11h12"))),
		html.Child(html.SvgPath(html.AD("M7 6h3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
