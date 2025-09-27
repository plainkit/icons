package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUp01 creates a Arrow Up 0 1 Lucide icon.
func ArrowUp01(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-0-1", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m3 8 4-4 4 4"))),
		html.Child(html.SvgPath(html.AD("M7 4v16"))),
		html.Child(html.SvgRect(html.AWidth("4"), html.AHeight("6"), html.AX("15"), html.AY("4"), html.ARy("2"))),
		html.Child(html.SvgPath(html.AD("M17 20v-6h-2"))),
		html.Child(html.SvgPath(html.AD("M15 20h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
