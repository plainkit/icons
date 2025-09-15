package lucide

import x "github.com/plainkit/html"

// HandCoins creates a Hand Coins Lucide icon.
func HandCoins(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-hand-coins", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 15h2a2 2 0 1 0 0-4h-3c-.6 0-1.1.2-1.4.6L3 17"))),
		x.Child(x.Path(x.D("m7 21 1.6-1.4c.3-.4.8-.6 1.4-.6h4c1.1 0 2.1-.4 2.8-1.2l4.6-4.4a2 2 0 0 0-2.75-2.91l-4.2 3.9"))),
		x.Child(x.Path(x.D("m2 16 6 6"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("9"), x.R("2.9"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("5"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
