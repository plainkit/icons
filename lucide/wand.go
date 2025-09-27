package lucide

import (
	html "github.com/plainkit/html"
)

// Wand creates a Wand Lucide icon.
func Wand(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wand", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 4V2"))),
		html.Child(html.SvgPath(html.AD("M15 16v-2"))),
		html.Child(html.SvgPath(html.AD("M8 9h2"))),
		html.Child(html.SvgPath(html.AD("M20 9h2"))),
		html.Child(html.SvgPath(html.AD("M17.8 11.8 19 13"))),
		html.Child(html.SvgPath(html.AD("M15 9h.01"))),
		html.Child(html.SvgPath(html.AD("M17.8 6.2 19 5"))),
		html.Child(html.SvgPath(html.AD("m3 21 9-9"))),
		html.Child(html.SvgPath(html.AD("M12.2 6.2 11 5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
