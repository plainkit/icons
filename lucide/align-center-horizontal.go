package lucide

import (
	html "github.com/plainkit/html"
)

// AlignCenterHorizontal creates a Align Center Horizontal Lucide icon.
func AlignCenterHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-center-horizontal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 12h20"))),
		html.Child(html.SvgPath(html.AD("M10 16v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-4"))),
		html.Child(html.SvgPath(html.AD("M10 8V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v4"))),
		html.Child(html.SvgPath(html.AD("M20 16v1a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-1"))),
		html.Child(html.SvgPath(html.AD("M14 8V7c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
