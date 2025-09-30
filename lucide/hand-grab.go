package lucide

import (
	html "github.com/plainkit/html"
)

// HandGrab creates a Hand Grab Lucide icon.
func HandGrab(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hand-grab", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18 11.5V9a2 2 0 0 0-2-2a2 2 0 0 0-2 2v1.4")),
		html.SvgPath(html.AD("M14 10V8a2 2 0 0 0-2-2a2 2 0 0 0-2 2v2")),
		html.SvgPath(html.AD("M10 9.9V9a2 2 0 0 0-2-2a2 2 0 0 0-2 2v5")),
		html.SvgPath(html.AD("M6 14a2 2 0 0 0-2-2a2 2 0 0 0-2 2")),
		html.SvgPath(html.AD("M18 11a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-4a8 8 0 0 1-8-8 2 2 0 1 1 4 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
