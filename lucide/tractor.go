package lucide

import (
	html "github.com/plainkit/html"
)

// Tractor creates a Tractor Lucide icon.
func Tractor(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tractor", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m10 11 11 .9a1 1 0 0 1 .8 1.1l-.665 4.158a1 1 0 0 1-.988.842H20"))),
		html.Child(html.SvgPath(html.AD("M16 18h-5"))),
		html.Child(html.SvgPath(html.AD("M18 5a1 1 0 0 0-1 1v5.573"))),
		html.Child(html.SvgPath(html.AD("M3 4h8.129a1 1 0 0 1 .99.863L13 11.246"))),
		html.Child(html.SvgPath(html.AD("M4 11V4"))),
		html.Child(html.SvgPath(html.AD("M7 15h.01"))),
		html.Child(html.SvgPath(html.AD("M8 10.1V4"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("2"))),
		html.Child(html.SvgCircle(html.ACx("7"), html.ACy("15"), html.AR("5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
