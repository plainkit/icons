package lucide

import (
	html "github.com/plainkit/html"
)

// Hexagon creates a Hexagon Lucide icon.
func Hexagon(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hexagon", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
