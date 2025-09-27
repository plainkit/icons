package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDown10 creates a Arrow Down 1 0 Lucide icon.
func ArrowDown10(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-1-0", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m3 16 4 4 4-4"))),
		html.Child(html.SvgPath(html.AD("M7 20V4"))),
		html.Child(html.SvgPath(html.AD("M17 10V4h-2"))),
		html.Child(html.SvgPath(html.AD("M15 10h4"))),
		html.Child(html.SvgRect(html.AWidth("4"), html.AHeight("6"), html.AX("15"), html.AY("14"), html.ARy("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
