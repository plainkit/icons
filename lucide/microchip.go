package lucide

import (
	html "github.com/plainkit/html"
)

// Microchip creates a Microchip Lucide icon.
func Microchip(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-microchip", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 12h2"))),
		html.Child(html.SvgPath(html.AD("M18 16h2"))),
		html.Child(html.SvgPath(html.AD("M18 20h2"))),
		html.Child(html.SvgPath(html.AD("M18 4h2"))),
		html.Child(html.SvgPath(html.AD("M18 8h2"))),
		html.Child(html.SvgPath(html.AD("M4 12h2"))),
		html.Child(html.SvgPath(html.AD("M4 16h2"))),
		html.Child(html.SvgPath(html.AD("M4 20h2"))),
		html.Child(html.SvgPath(html.AD("M4 4h2"))),
		html.Child(html.SvgPath(html.AD("M4 8h2"))),
		html.Child(html.SvgPath(html.AD("M8 2a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2h-1.5c-.276 0-.494.227-.562.495a2 2 0 0 1-3.876 0C9.994 2.227 9.776 2 9.5 2z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
