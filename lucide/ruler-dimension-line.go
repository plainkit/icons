package lucide

import (
	html "github.com/plainkit/html"
)

// RulerDimensionLine creates a Ruler Dimension Line Lucide icon.
func RulerDimensionLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ruler-dimension-line", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 15v-3.014"))),
		html.Child(html.SvgPath(html.AD("M16 15v-3.014"))),
		html.Child(html.SvgPath(html.AD("M20 6H4"))),
		html.Child(html.SvgPath(html.AD("M20 8V4"))),
		html.Child(html.SvgPath(html.AD("M4 8V4"))),
		html.Child(html.SvgPath(html.AD("M8 15v-3.014"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("7"), html.AX("3"), html.AY("12"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
