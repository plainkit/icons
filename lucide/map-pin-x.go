package lucide

import (
	html "github.com/plainkit/html"
)

// MapPinX creates a Map Pin X Lucide icon.
func MapPinX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-pin-x", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M19.752 11.901A7.78 7.78 0 0 0 20 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 19 19 0 0 0 .09-.077"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("m21.5 15.5-5 5"))),
		html.Child(html.SvgPath(html.AD("m21.5 20.5-5-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
