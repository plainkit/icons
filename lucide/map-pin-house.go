package lucide

import (
	html "github.com/plainkit/html"
)

// MapPinHouse creates a Map Pin House Lucide icon.
func MapPinHouse(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-pin-house", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 22a1 1 0 0 1-1-1v-4a1 1 0 0 1 .445-.832l3-2a1 1 0 0 1 1.11 0l3 2A1 1 0 0 1 22 17v4a1 1 0 0 1-1 1z"))),
		html.Child(html.SvgPath(html.AD("M18 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 .601.2"))),
		html.Child(html.SvgPath(html.AD("M18 22v-3"))),
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("10"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
