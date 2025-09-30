package lucide

import (
	html "github.com/plainkit/html"
)

// PencilRuler creates a Pencil Ruler Lucide icon.
func PencilRuler(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pencil-ruler", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13 7 8.7 2.7a2.41 2.41 0 0 0-3.4 0L2.7 5.3a2.41 2.41 0 0 0 0 3.4L7 13")),
		html.SvgPath(html.AD("m8 6 2-2")),
		html.SvgPath(html.AD("m18 16 2-2")),
		html.SvgPath(html.AD("m17 11 4.3 4.3c.94.94.94 2.46 0 3.4l-2.6 2.6c-.94.94-2.46.94-3.4 0L11 17")),
		html.SvgPath(html.AD("M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z")),
		html.SvgPath(html.AD("m15 5 4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
