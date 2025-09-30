package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUpZA creates a Arrow Up Z A Lucide icon.
func ArrowUpZA(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-z-a", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m3 8 4-4 4 4")),
		html.SvgPath(html.AD("M7 4v16")),
		html.SvgPath(html.AD("M15 4h5l-5 6h5")),
		html.SvgPath(html.AD("M15 20v-3.5a2.5 2.5 0 0 1 5 0V20")),
		html.SvgPath(html.AD("M20 18h-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
