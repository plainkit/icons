package lucide

import (
	html "github.com/plainkit/html"
)

// ShoppingBag creates a Shopping Bag Lucide icon.
func ShoppingBag(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shopping-bag", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 10a4 4 0 0 1-8 0"))),
		html.Child(html.SvgPath(html.AD("M3.103 6.034h17.794"))),
		html.Child(html.SvgPath(html.AD("M3.4 5.467a2 2 0 0 0-.4 1.2V20a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6.667a2 2 0 0 0-.4-1.2l-2-2.667A2 2 0 0 0 17 2H7a2 2 0 0 0-1.6.8z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
