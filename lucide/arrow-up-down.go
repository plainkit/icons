package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUpDown creates a Arrow Up Down Lucide icon.
func ArrowUpDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m21 16-4 4-4-4")),
		html.SvgPath(html.AD("M17 20V4")),
		html.SvgPath(html.AD("m3 8 4-4 4 4")),
		html.SvgPath(html.AD("M7 4v16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
