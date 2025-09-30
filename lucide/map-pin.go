package lucide

import (
	html "github.com/plainkit/html"
)

// MapPin creates a Map Pin Lucide icon.
func MapPin(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-pin", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0")),
		html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
