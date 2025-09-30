package lucide

import (
	html "github.com/plainkit/html"
)

// MapPinPlus creates a Map Pin Plus Lucide icon.
func MapPinPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-pin-plus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M19.914 11.105A7.298 7.298 0 0 0 20 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 32 32 0 0 0 .824-.738")),
		html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("3")),
		html.SvgPath(html.AD("M16 18h6")),
		html.SvgPath(html.AD("M19 15v6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
