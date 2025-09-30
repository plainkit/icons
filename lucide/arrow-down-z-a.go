package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDownZA creates a Arrow Down Z A Lucide icon.
func ArrowDownZA(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-z-a", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m3 16 4 4 4-4")),
		html.SvgPath(html.AD("M7 4v16")),
		html.SvgPath(html.AD("M15 4h5l-5 6h5")),
		html.SvgPath(html.AD("M15 20v-3.5a2.5 2.5 0 0 1 5 0V20")),
		html.SvgPath(html.AD("M20 18h-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
