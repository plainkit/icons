package lucide

import (
	html "github.com/plainkit/html"
)

// MapPinCheckInside creates a Map Pin Check Inside Lucide icon.
func MapPinCheckInside(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-pin-check-inside", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0"))),
		html.Child(html.SvgPath(html.AD("m9 10 2 2 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
