package lucide

import (
	html "github.com/plainkit/html"
)

// ChartArea creates a Chart Area Lucide icon.
func ChartArea(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-area", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16")),
		html.SvgPath(html.AD("M7 11.207a.5.5 0 0 1 .146-.353l2-2a.5.5 0 0 1 .708 0l3.292 3.292a.5.5 0 0 0 .708 0l4.292-4.292a.5.5 0 0 1 .854.353V16a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
