package lucide

import (
	html "github.com/plainkit/html"
)

// ChartBarBig creates a Chart Bar Big Lucide icon.
func ChartBarBig(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-bar-big", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16")),
		html.SvgRect(html.AWidth("9"), html.AHeight("4"), html.AX("7"), html.AY("13"), html.ARx("1")),
		html.SvgRect(html.AWidth("12"), html.AHeight("4"), html.AX("7"), html.AY("5"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
