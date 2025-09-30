package lucide

import (
	html "github.com/plainkit/html"
)

// ClosedCaption creates a Closed Caption Lucide icon.
func ClosedCaption(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-closed-caption", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 9.17a3 3 0 1 0 0 5.66")),
		html.SvgPath(html.AD("M17 9.17a3 3 0 1 0 0 5.66")),
		html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("5"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
