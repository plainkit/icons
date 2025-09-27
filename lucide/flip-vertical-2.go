package lucide

import (
	html "github.com/plainkit/html"
)

// FlipVertical2 creates a Flip Vertical 2 Lucide icon.
func FlipVertical2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flip-vertical-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m17 3-5 5-5-5h10"))),
		html.Child(html.SvgPath(html.AD("m17 21-5-5-5 5h10"))),
		html.Child(html.SvgPath(html.AD("M4 12H2"))),
		html.Child(html.SvgPath(html.AD("M10 12H8"))),
		html.Child(html.SvgPath(html.AD("M16 12h-2"))),
		html.Child(html.SvgPath(html.AD("M22 12h-2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
