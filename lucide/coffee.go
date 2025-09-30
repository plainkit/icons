package lucide

import (
	html "github.com/plainkit/html"
)

// Coffee creates a Coffee Lucide icon.
func Coffee(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-coffee", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 2v2")),
		html.SvgPath(html.AD("M14 2v2")),
		html.SvgPath(html.AD("M16 8a1 1 0 0 1 1 1v8a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V9a1 1 0 0 1 1-1h14a4 4 0 1 1 0 8h-1")),
		html.SvgPath(html.AD("M6 2v2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
