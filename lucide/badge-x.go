package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeX creates a Badge X Lucide icon.
func BadgeX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z")),
		html.SvgLine(html.AX1("15"), html.AX2("9"), html.AY1("9"), html.AY2("15")),
		html.SvgLine(html.AX1("9"), html.AX2("15"), html.AY1("9"), html.AY2("15")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
