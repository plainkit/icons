package lucide

import (
	html "github.com/plainkit/html"
)

// ChartLine creates a Chart Line Lucide icon.
func ChartLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-line", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16"))),
		html.Child(html.SvgPath(html.AD("m19 9-5 5-4-4-3 3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
