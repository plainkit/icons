package lucide

import (
	html "github.com/plainkit/html"
)

// TramFront creates a Tram Front Lucide icon.
func TramFront(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tram-front", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("16"), html.AHeight("16"), html.AX("4"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M4 11h16")),
		html.SvgPath(html.AD("M12 3v8")),
		html.SvgPath(html.AD("m8 19-2 3")),
		html.SvgPath(html.AD("m18 22-2-3")),
		html.SvgPath(html.AD("M8 15h.01")),
		html.SvgPath(html.AD("M16 15h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
