package lucide

import (
	html "github.com/plainkit/html"
)

// MapPinCheck creates a Map Pin Check Lucide icon.
func MapPinCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-pin-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M19.43 12.935c.357-.967.57-1.955.57-2.935a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 32.197 32.197 0 0 0 .813-.728")),
		html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("3")),
		html.SvgPath(html.AD("m16 18 2 2 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
