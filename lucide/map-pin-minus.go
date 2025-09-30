package lucide

import (
	html "github.com/plainkit/html"
)

// MapPinMinus creates a Map Pin Minus Lucide icon.
func MapPinMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-pin-minus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18.977 14C19.6 12.701 20 11.343 20 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 32 32 0 0 0 .824-.738")),
		html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("3")),
		html.SvgPath(html.AD("M16 18h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
