package lucide

import (
	html "github.com/plainkit/html"
)

// Building creates a Building Lucide icon.
func Building(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-building", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 10h.01")),
		html.SvgPath(html.AD("M12 14h.01")),
		html.SvgPath(html.AD("M12 6h.01")),
		html.SvgPath(html.AD("M16 10h.01")),
		html.SvgPath(html.AD("M16 14h.01")),
		html.SvgPath(html.AD("M16 6h.01")),
		html.SvgPath(html.AD("M8 10h.01")),
		html.SvgPath(html.AD("M8 14h.01")),
		html.SvgPath(html.AD("M8 6h.01")),
		html.SvgPath(html.AD("M9 22v-3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v3")),
		html.SvgRect(html.AWidth("16"), html.AHeight("20"), html.AX("4"), html.AY("2"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
