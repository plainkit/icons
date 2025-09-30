package lucide

import (
	html "github.com/plainkit/html"
)

// ChartColumnStacked creates a Chart Column Stacked Lucide icon.
func ChartColumnStacked(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-column-stacked", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 13H7")),
		html.SvgPath(html.AD("M19 9h-4")),
		html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16")),
		html.SvgRect(html.AWidth("4"), html.AHeight("12"), html.AX("15"), html.AY("5"), html.ARx("1")),
		html.SvgRect(html.AWidth("4"), html.AHeight("9"), html.AX("7"), html.AY("8"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
