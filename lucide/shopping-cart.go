package lucide

import (
	html "github.com/plainkit/html"
)

// ShoppingCart creates a Shopping Cart Lucide icon.
func ShoppingCart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shopping-cart", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("8"), html.ACy("21"), html.AR("1"))),
		html.Child(html.SvgCircle(html.ACx("19"), html.ACy("21"), html.AR("1"))),
		html.Child(html.SvgPath(html.AD("M2.05 2.05h2l2.66 12.42a2 2 0 0 0 2 1.58h9.78a2 2 0 0 0 1.95-1.57l1.65-7.43H5.12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
