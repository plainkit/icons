package lucide

import (
	html "github.com/plainkit/html"
)

// CarTaxiFront creates a Car Taxi Front Lucide icon.
func CarTaxiFront(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-car-taxi-front", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 2h4")),
		html.SvgPath(html.AD("m21 8-2 2-1.5-3.7A2 2 0 0 0 15.646 5H8.4a2 2 0 0 0-1.903 1.257L5 10 3 8")),
		html.SvgPath(html.AD("M7 14h.01")),
		html.SvgPath(html.AD("M17 14h.01")),
		html.SvgRect(html.AWidth("18"), html.AHeight("8"), html.AX("3"), html.AY("10"), html.ARx("2")),
		html.SvgPath(html.AD("M5 18v2")),
		html.SvgPath(html.AD("M19 18v2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
