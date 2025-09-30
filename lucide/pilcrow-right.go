package lucide

import (
	html "github.com/plainkit/html"
)

// PilcrowRight creates a Pilcrow Right Lucide icon.
func PilcrowRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pilcrow-right", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 3v11")),
		html.SvgPath(html.AD("M10 9H7a1 1 0 0 1 0-6h8")),
		html.SvgPath(html.AD("M14 3v11")),
		html.SvgPath(html.AD("m18 14 4 4H2")),
		html.SvgPath(html.AD("m22 18-4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
