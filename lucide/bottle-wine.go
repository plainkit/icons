package lucide

import (
	html "github.com/plainkit/html"
)

// BottleWine creates a Bottle Wine Lucide icon.
func BottleWine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bottle-wine", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a6 6 0 0 0 1.2 3.6l.6.8A6 6 0 0 1 17 13v8a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-8a6 6 0 0 1 1.2-3.6l.6-.8A6 6 0 0 0 10 5z")),
		html.SvgPath(html.AD("M17 13h-4a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
