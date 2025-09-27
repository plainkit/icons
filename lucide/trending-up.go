package lucide

import (
	html "github.com/plainkit/html"
)

// TrendingUp creates a Trending Up Lucide icon.
func TrendingUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-trending-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 7h6v6"))),
		html.Child(html.SvgPath(html.AD("m22 7-8.5 8.5-5-5L2 17"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
