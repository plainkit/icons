package lucide

import (
	html "github.com/plainkit/html"
)

// Heater creates a Heater Lucide icon.
func Heater(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-heater", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 8c2-3-2-3 0-6"))),
		html.Child(html.SvgPath(html.AD("M15.5 8c2-3-2-3 0-6"))),
		html.Child(html.SvgPath(html.AD("M6 10h.01"))),
		html.Child(html.SvgPath(html.AD("M6 14h.01"))),
		html.Child(html.SvgPath(html.AD("M10 16v-4"))),
		html.Child(html.SvgPath(html.AD("M14 16v-4"))),
		html.Child(html.SvgPath(html.AD("M18 16v-4"))),
		html.Child(html.SvgPath(html.AD("M20 6a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h3"))),
		html.Child(html.SvgPath(html.AD("M5 20v2"))),
		html.Child(html.SvgPath(html.AD("M19 20v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
