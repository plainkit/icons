package lucide

import (
	html "github.com/plainkit/html"
)

// Rabbit creates a Rabbit Lucide icon.
func Rabbit(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rabbit", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13 16a3 3 0 0 1 2.24 5"))),
		html.Child(html.SvgPath(html.AD("M18 12h.01"))),
		html.Child(html.SvgPath(html.AD("M18 21h-8a4 4 0 0 1-4-4 7 7 0 0 1 7-7h.2L9.6 6.4a1 1 0 1 1 2.8-2.8L15.8 7h.2c3.3 0 6 2.7 6 6v1a2 2 0 0 1-2 2h-1a3 3 0 0 0-3 3"))),
		html.Child(html.SvgPath(html.AD("M20 8.54V4a2 2 0 1 0-4 0v3"))),
		html.Child(html.SvgPath(html.AD("M7.612 12.524a3 3 0 1 0-1.6 4.3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
