package lucide

import (
	html "github.com/plainkit/html"
)

// FlipVertical creates a Flip Vertical Lucide icon.
func FlipVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flip-vertical", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 8V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v3"))),
		html.Child(html.SvgPath(html.AD("M21 16v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-3"))),
		html.Child(html.SvgPath(html.AD("M4 12H2"))),
		html.Child(html.SvgPath(html.AD("M10 12H8"))),
		html.Child(html.SvgPath(html.AD("M16 12h-2"))),
		html.Child(html.SvgPath(html.AD("M22 12h-2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
