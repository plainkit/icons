package lucide

import (
	html "github.com/plainkit/html"
)

// PaintbrushVertical creates a Paintbrush Vertical Lucide icon.
func PaintbrushVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-paintbrush-vertical", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 2v2"))),
		html.Child(html.SvgPath(html.AD("M14 2v4"))),
		html.Child(html.SvgPath(html.AD("M17 2a1 1 0 0 1 1 1v9H6V3a1 1 0 0 1 1-1z"))),
		html.Child(html.SvgPath(html.AD("M6 12a1 1 0 0 0-1 1v1a2 2 0 0 0 2 2h2a1 1 0 0 1 1 1v2.9a2 2 0 1 0 4 0V17a1 1 0 0 1 1-1h2a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
