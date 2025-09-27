package lucide

import (
	html "github.com/plainkit/html"
)

// Trash2 creates a Trash 2 Lucide icon.
func Trash2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-trash-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 11v6"))),
		html.Child(html.SvgPath(html.AD("M14 11v6"))),
		html.Child(html.SvgPath(html.AD("M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6"))),
		html.Child(html.SvgPath(html.AD("M3 6h18"))),
		html.Child(html.SvgPath(html.AD("M8 6V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
