package lucide

import (
	html "github.com/plainkit/html"
)

// UtensilsCrossed creates a Utensils Crossed Lucide icon.
func UtensilsCrossed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-utensils-crossed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m16 2-2.3 2.3a3 3 0 0 0 0 4.2l1.8 1.8a3 3 0 0 0 4.2 0L22 8"))),
		html.Child(html.SvgPath(html.AD("M15 15 3.3 3.3a4.2 4.2 0 0 0 0 6l7.3 7.3c.7.7 2 .7 2.8 0L15 15Zm0 0 7 7"))),
		html.Child(html.SvgPath(html.AD("m2.1 21.8 6.4-6.3"))),
		html.Child(html.SvgPath(html.AD("m19 5-7 7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
