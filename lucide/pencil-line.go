package lucide

import (
	html "github.com/plainkit/html"
)

// PencilLine creates a Pencil Line Lucide icon.
func PencilLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pencil-line", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13 21h8"))),
		html.Child(html.SvgPath(html.AD("m15 5 4 4"))),
		html.Child(html.SvgPath(html.AD("M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
