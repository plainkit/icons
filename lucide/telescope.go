package lucide

import (
	html "github.com/plainkit/html"
)

// Telescope creates a Telescope Lucide icon.
func Telescope(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-telescope", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m10.065 12.493-6.18 1.318a.934.934 0 0 1-1.108-.702l-.537-2.15a1.07 1.07 0 0 1 .691-1.265l13.504-4.44"))),
		html.Child(html.SvgPath(html.AD("m13.56 11.747 4.332-.924"))),
		html.Child(html.SvgPath(html.AD("m16 21-3.105-6.21"))),
		html.Child(html.SvgPath(html.AD("M16.485 5.94a2 2 0 0 1 1.455-2.425l1.09-.272a1 1 0 0 1 1.212.727l1.515 6.06a1 1 0 0 1-.727 1.213l-1.09.272a2 2 0 0 1-2.425-1.455z"))),
		html.Child(html.SvgPath(html.AD("m6.158 8.633 1.114 4.456"))),
		html.Child(html.SvgPath(html.AD("m8 21 3.105-6.21"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("13"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
