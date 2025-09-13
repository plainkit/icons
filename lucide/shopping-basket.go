package lucide

import x "github.com/bloxui/blox"

// ShoppingBasket creates a Shopping Basket Lucide icon.
func ShoppingBasket(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-shopping-basket", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 11-1 9"))),
		x.Child(x.Path(x.D("m19 11-4-7"))),
		x.Child(x.Path(x.D("M2 11h20"))),
		x.Child(x.Path(x.D("m3.5 11 1.6 7.4a2 2 0 0 0 2 1.6h9.8a2 2 0 0 0 2-1.6l1.7-7.4"))),
		x.Child(x.Path(x.D("M4.5 15.5h15"))),
		x.Child(x.Path(x.D("m5 11 4-7"))),
		x.Child(x.Path(x.D("m9 11 1 9"))),
	)
	return x.Svg(svgArgs...)
}
