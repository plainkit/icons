package lucide

import (
	html "github.com/plainkit/html"
)

// Ruler creates a Ruler Lucide icon.
func Ruler(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ruler", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21.3 15.3a2.4 2.4 0 0 1 0 3.4l-2.6 2.6a2.4 2.4 0 0 1-3.4 0L2.7 8.7a2.41 2.41 0 0 1 0-3.4l2.6-2.6a2.41 2.41 0 0 1 3.4 0Z"))),
		html.Child(html.SvgPath(html.AD("m14.5 12.5 2-2"))),
		html.Child(html.SvgPath(html.AD("m11.5 9.5 2-2"))),
		html.Child(html.SvgPath(html.AD("m8.5 6.5 2-2"))),
		html.Child(html.SvgPath(html.AD("m17.5 15.5 2-2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
