package lucide

import (
	html "github.com/plainkit/html"
)

// Angry creates a Angry Lucide icon.
func Angry(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-angry", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("M16 16s-1.5-2-4-2-4 2-4 2"))),
		html.Child(html.SvgPath(html.AD("M7.5 8 10 9"))),
		html.Child(html.SvgPath(html.AD("m14 9 2.5-1"))),
		html.Child(html.SvgPath(html.AD("M9 10h.01"))),
		html.Child(html.SvgPath(html.AD("M15 10h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
