package lucide

import (
	html "github.com/plainkit/html"
)

// CupSoda creates a Cup Soda Lucide icon.
func CupSoda(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cup-soda", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m6 8 1.75 12.28a2 2 0 0 0 2 1.72h4.54a2 2 0 0 0 2-1.72L18 8")),
		html.SvgPath(html.AD("M5 8h14")),
		html.SvgPath(html.AD("M7 15a6.47 6.47 0 0 1 5 0 6.47 6.47 0 0 0 5 0")),
		html.SvgPath(html.AD("m12 8 1-6h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
