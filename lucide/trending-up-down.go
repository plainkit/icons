package lucide

import (
	html "github.com/plainkit/html"
)

// TrendingUpDown creates a Trending Up Down Lucide icon.
func TrendingUpDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-trending-up-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14.828 14.828 21 21")),
		html.SvgPath(html.AD("M21 16v5h-5")),
		html.SvgPath(html.AD("m21 3-9 9-4-4-6 6")),
		html.SvgPath(html.AD("M21 8V3h-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
