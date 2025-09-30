package lucide

import (
	html "github.com/plainkit/html"
)

// Binoculars creates a Binoculars Lucide icon.
func Binoculars(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-binoculars", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 10h4")),
		html.SvgPath(html.AD("M19 7V4a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v3")),
		html.SvgPath(html.AD("M20 21a2 2 0 0 0 2-2v-3.851c0-1.39-2-2.962-2-4.829V8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v11a2 2 0 0 0 2 2z")),
		html.SvgPath(html.AD("M 22 16 L 2 16")),
		html.SvgPath(html.AD("M4 21a2 2 0 0 1-2-2v-3.851c0-1.39 2-2.962 2-4.829V8a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v11a2 2 0 0 1-2 2z")),
		html.SvgPath(html.AD("M9 7V4a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
