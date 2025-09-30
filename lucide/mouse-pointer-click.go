package lucide

import (
	html "github.com/plainkit/html"
)

// MousePointerClick creates a Mouse Pointer Click Lucide icon.
func MousePointerClick(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mouse-pointer-click", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 4.1 12 6")),
		html.SvgPath(html.AD("m5.1 8-2.9-.8")),
		html.SvgPath(html.AD("m6 12-1.9 2")),
		html.SvgPath(html.AD("M7.2 2.2 8 5.1")),
		html.SvgPath(html.AD("M9.037 9.69a.498.498 0 0 1 .653-.653l11 4.5a.5.5 0 0 1-.074.949l-4.349 1.041a1 1 0 0 0-.74.739l-1.04 4.35a.5.5 0 0 1-.95.074z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
