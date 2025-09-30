package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowRightLeft creates a Arrow Right Left Lucide icon.
func ArrowRightLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-right-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m16 3 4 4-4 4")),
		html.SvgPath(html.AD("M20 7H4")),
		html.SvgPath(html.AD("m8 21-4-4 4-4")),
		html.SvgPath(html.AD("M4 17h16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
