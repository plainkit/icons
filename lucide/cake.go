package lucide

import (
	html "github.com/plainkit/html"
)

// Cake creates a Cake Lucide icon.
func Cake(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cake", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20 21v-8a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v8"))),
		html.Child(html.SvgPath(html.AD("M4 16s.5-1 2-1 2.5 2 4 2 2.5-2 4-2 2.5 2 4 2 2-1 2-1"))),
		html.Child(html.SvgPath(html.AD("M2 21h20"))),
		html.Child(html.SvgPath(html.AD("M7 8v3"))),
		html.Child(html.SvgPath(html.AD("M12 8v3"))),
		html.Child(html.SvgPath(html.AD("M17 8v3"))),
		html.Child(html.SvgPath(html.AD("M7 4h.01"))),
		html.Child(html.SvgPath(html.AD("M12 4h.01"))),
		html.Child(html.SvgPath(html.AD("M17 4h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
