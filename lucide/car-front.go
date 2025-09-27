package lucide

import (
	html "github.com/plainkit/html"
)

// CarFront creates a Car Front Lucide icon.
func CarFront(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-car-front", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m21 8-2 2-1.5-3.7A2 2 0 0 0 15.646 5H8.4a2 2 0 0 0-1.903 1.257L5 10 3 8"))),
		html.Child(html.SvgPath(html.AD("M7 14h.01"))),
		html.Child(html.SvgPath(html.AD("M17 14h.01"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("8"), html.AX("3"), html.AY("10"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M5 18v2"))),
		html.Child(html.SvgPath(html.AD("M19 18v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
