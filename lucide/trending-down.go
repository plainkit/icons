package lucide

import (
	html "github.com/plainkit/html"
)

// TrendingDown creates a Trending Down Lucide icon.
func TrendingDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-trending-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 17h6v-6")),
		html.SvgPath(html.AD("m22 17-8.5-8.5-5 5L2 7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
