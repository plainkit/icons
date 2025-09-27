package lucide

import (
	html "github.com/plainkit/html"
)

// Bug creates a Bug Lucide icon.
func Bug(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bug", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 20v-9"))),
		html.Child(html.SvgPath(html.AD("M14 7a4 4 0 0 1 4 4v3a6 6 0 0 1-12 0v-3a4 4 0 0 1 4-4z"))),
		html.Child(html.SvgPath(html.AD("M14.12 3.88 16 2"))),
		html.Child(html.SvgPath(html.AD("M21 21a4 4 0 0 0-3.81-4"))),
		html.Child(html.SvgPath(html.AD("M21 5a4 4 0 0 1-3.55 3.97"))),
		html.Child(html.SvgPath(html.AD("M22 13h-4"))),
		html.Child(html.SvgPath(html.AD("M3 21a4 4 0 0 1 3.81-4"))),
		html.Child(html.SvgPath(html.AD("M3 5a4 4 0 0 0 3.55 3.97"))),
		html.Child(html.SvgPath(html.AD("M6 13H2"))),
		html.Child(html.SvgPath(html.AD("m8 2 1.88 1.88"))),
		html.Child(html.SvgPath(html.AD("M9 7.13V6a3 3 0 1 1 6 0v1.13"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
