package lucide

import (
	html "github.com/plainkit/html"
)

// Hotel creates a Hotel Lucide icon.
func Hotel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hotel", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 22v-6.57")),
		html.SvgPath(html.AD("M12 11h.01")),
		html.SvgPath(html.AD("M12 7h.01")),
		html.SvgPath(html.AD("M14 15.43V22")),
		html.SvgPath(html.AD("M15 16a5 5 0 0 0-6 0")),
		html.SvgPath(html.AD("M16 11h.01")),
		html.SvgPath(html.AD("M16 7h.01")),
		html.SvgPath(html.AD("M8 11h.01")),
		html.SvgPath(html.AD("M8 7h.01")),
		html.SvgRect(html.AWidth("16"), html.AHeight("20"), html.AX("4"), html.AY("2"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
