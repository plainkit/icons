package lucide

import (
	html "github.com/plainkit/html"
)

// MapPinPlusInside creates a Map Pin Plus Inside Lucide icon.
func MapPinPlusInside(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-pin-plus-inside", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0")),
		html.SvgPath(html.AD("M12 7v6")),
		html.SvgPath(html.AD("M9 10h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
