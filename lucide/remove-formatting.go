package lucide

import (
	html "github.com/plainkit/html"
)

// RemoveFormatting creates a Remove Formatting Lucide icon.
func RemoveFormatting(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-remove-formatting", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 7V4h16v3")),
		html.SvgPath(html.AD("M5 20h6")),
		html.SvgPath(html.AD("M13 4 8 20")),
		html.SvgPath(html.AD("m15 15 5 5")),
		html.SvgPath(html.AD("m20 15-5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
