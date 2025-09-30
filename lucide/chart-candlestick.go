package lucide

import (
	html "github.com/plainkit/html"
)

// ChartCandlestick creates a Chart Candlestick Lucide icon.
func ChartCandlestick(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-candlestick", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M9 5v4")),
		html.SvgRect(html.AWidth("4"), html.AHeight("6"), html.AX("7"), html.AY("9"), html.ARx("1")),
		html.SvgPath(html.AD("M9 15v2")),
		html.SvgPath(html.AD("M17 3v2")),
		html.SvgRect(html.AWidth("4"), html.AHeight("8"), html.AX("15"), html.AY("5"), html.ARx("1")),
		html.SvgPath(html.AD("M17 13v3")),
		html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
