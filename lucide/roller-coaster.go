package lucide

import (
	html "github.com/plainkit/html"
)

// RollerCoaster creates a Roller Coaster Lucide icon.
func RollerCoaster(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-roller-coaster", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 19V5"))),
		html.Child(html.SvgPath(html.AD("M10 19V6.8"))),
		html.Child(html.SvgPath(html.AD("M14 19v-7.8"))),
		html.Child(html.SvgPath(html.AD("M18 5v4"))),
		html.Child(html.SvgPath(html.AD("M18 19v-6"))),
		html.Child(html.SvgPath(html.AD("M22 19V9"))),
		html.Child(html.SvgPath(html.AD("M2 19V9a4 4 0 0 1 4-4c2 0 4 1.33 6 4s4 4 6 4a4 4 0 1 0-3-6.65"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
