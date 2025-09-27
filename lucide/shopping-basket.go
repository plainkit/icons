package lucide

import (
	html "github.com/plainkit/html"
)

// ShoppingBasket creates a Shopping Basket Lucide icon.
func ShoppingBasket(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shopping-basket", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m15 11-1 9"))),
		html.Child(html.SvgPath(html.AD("m19 11-4-7"))),
		html.Child(html.SvgPath(html.AD("M2 11h20"))),
		html.Child(html.SvgPath(html.AD("m3.5 11 1.6 7.4a2 2 0 0 0 2 1.6h9.8a2 2 0 0 0 2-1.6l1.7-7.4"))),
		html.Child(html.SvgPath(html.AD("M4.5 15.5h15"))),
		html.Child(html.SvgPath(html.AD("m5 11 4-7"))),
		html.Child(html.SvgPath(html.AD("m9 11 1 9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
