package lucide

import (
	html "github.com/plainkit/html"
)

// PilcrowLeft creates a Pilcrow Left Lucide icon.
func PilcrowLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pilcrow-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 3v11")),
		html.SvgPath(html.AD("M14 9h-3a3 3 0 0 1 0-6h9")),
		html.SvgPath(html.AD("M18 3v11")),
		html.SvgPath(html.AD("M22 18H2l4-4")),
		html.SvgPath(html.AD("m6 22-4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
