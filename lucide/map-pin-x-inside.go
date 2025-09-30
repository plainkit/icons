package lucide

import (
	html "github.com/plainkit/html"
)

// MapPinXInside creates a Map Pin X Inside Lucide icon.
func MapPinXInside(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-pin-x-inside", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0")),
		html.SvgPath(html.AD("m14.5 7.5-5 5")),
		html.SvgPath(html.AD("m9.5 7.5 5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
