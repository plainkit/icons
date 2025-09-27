package lucide

import (
	html "github.com/plainkit/html"
)

// Fullscreen creates a Fullscreen Lucide icon.
func Fullscreen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-fullscreen", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 7V5a2 2 0 0 1 2-2h2"))),
		html.Child(html.SvgPath(html.AD("M17 3h2a2 2 0 0 1 2 2v2"))),
		html.Child(html.SvgPath(html.AD("M21 17v2a2 2 0 0 1-2 2h-2"))),
		html.Child(html.SvgPath(html.AD("M7 21H5a2 2 0 0 1-2-2v-2"))),
		html.Child(html.SvgRect(html.AWidth("10"), html.AHeight("8"), html.AX("7"), html.AY("8"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
