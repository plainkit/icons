package lucide

import (
	html "github.com/plainkit/html"
)

// HandCoins creates a Hand Coins Lucide icon.
func HandCoins(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hand-coins", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 15h2a2 2 0 1 0 0-4h-3c-.6 0-1.1.2-1.4.6L3 17"))),
		html.Child(html.SvgPath(html.AD("m7 21 1.6-1.4c.3-.4.8-.6 1.4-.6h4c1.1 0 2.1-.4 2.8-1.2l4.6-4.4a2 2 0 0 0-2.75-2.91l-4.2 3.9"))),
		html.Child(html.SvgPath(html.AD("m2 16 6 6"))),
		html.Child(html.SvgCircle(html.ACx("16"), html.ACy("9"), html.AR("2.9"))),
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("5"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
