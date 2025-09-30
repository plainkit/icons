package lucide

import (
	html "github.com/plainkit/html"
)

// TriangleRight creates a Triangle Right Lucide icon.
func TriangleRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-triangle-right", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 18a2 2 0 0 1-2 2H3c-1.1 0-1.3-.6-.4-1.3L20.4 4.3c.9-.7 1.6-.4 1.6.7Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
