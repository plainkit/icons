package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUp10 creates a Arrow Up 1 0 Lucide icon.
func ArrowUp10(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-1-0", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m3 8 4-4 4 4"))),
		html.Child(html.SvgPath(html.AD("M7 4v16"))),
		html.Child(html.SvgPath(html.AD("M17 10V4h-2"))),
		html.Child(html.SvgPath(html.AD("M15 10h4"))),
		html.Child(html.SvgRect(html.AWidth("4"), html.AHeight("6"), html.AX("15"), html.AY("14"), html.ARy("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
