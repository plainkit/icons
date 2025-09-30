package lucide

import (
	html "github.com/plainkit/html"
)

// Bath creates a Bath Lucide icon.
func Bath(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bath", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 4 8 6")),
		html.SvgPath(html.AD("M17 19v2")),
		html.SvgPath(html.AD("M2 12h20")),
		html.SvgPath(html.AD("M7 19v2")),
		html.SvgPath(html.AD("M9 5 7.621 3.621A2.121 2.121 0 0 0 4 5v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
