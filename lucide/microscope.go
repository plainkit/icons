package lucide

import (
	html "github.com/plainkit/html"
)

// Microscope creates a Microscope Lucide icon.
func Microscope(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-microscope", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 18h8"))),
		html.Child(html.SvgPath(html.AD("M3 22h18"))),
		html.Child(html.SvgPath(html.AD("M14 22a7 7 0 1 0 0-14h-1"))),
		html.Child(html.SvgPath(html.AD("M9 14h2"))),
		html.Child(html.SvgPath(html.AD("M9 12a2 2 0 0 1-2-2V6h6v4a2 2 0 0 1-2 2Z"))),
		html.Child(html.SvgPath(html.AD("M12 6V3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
