package lucide

import (
	html "github.com/plainkit/html"
)

// ChartColumnBig creates a Chart Column Big Lucide icon.
func ChartColumnBig(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-column-big", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16"))),
		html.Child(html.SvgRect(html.AWidth("4"), html.AHeight("12"), html.AX("15"), html.AY("5"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("4"), html.AHeight("9"), html.AX("7"), html.AY("8"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
