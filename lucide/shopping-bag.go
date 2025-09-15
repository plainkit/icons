package lucide

import x "github.com/plainkit/html"

// ShoppingBag creates a Shopping Bag Lucide icon.
func ShoppingBag(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shopping-bag", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 10a4 4 0 0 1-8 0"))),
		x.Child(x.Path(x.D("M3.103 6.034h17.794"))),
		x.Child(x.Path(x.D("M3.4 5.467a2 2 0 0 0-.4 1.2V20a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6.667a2 2 0 0 0-.4-1.2l-2-2.667A2 2 0 0 0 17 2H7a2 2 0 0 0-1.6.8z"))),
	)
	return x.Svg(svgArgs...)
}
