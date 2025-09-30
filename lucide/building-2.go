package lucide

import (
	html "github.com/plainkit/html"
)

// Building2 creates a Building 2 Lucide icon.
func Building2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-building-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 22V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v18Z")),
		html.SvgPath(html.AD("M6 12H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2")),
		html.SvgPath(html.AD("M18 9h2a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2h-2")),
		html.SvgPath(html.AD("M10 6h4")),
		html.SvgPath(html.AD("M10 10h4")),
		html.SvgPath(html.AD("M10 14h4")),
		html.SvgPath(html.AD("M10 18h4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
