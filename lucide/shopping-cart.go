package lucide

import x "github.com/plainkit/blox"

// ShoppingCart creates a Shopping Cart Lucide icon.
func ShoppingCart(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shopping-cart", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("8"), x.Cy("21"), x.R("1"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("21"), x.R("1"))),
		x.Child(x.Path(x.D("M2.05 2.05h2l2.66 12.42a2 2 0 0 0 2 1.58h9.78a2 2 0 0 0 1.95-1.57l1.65-7.43H5.12"))),
	)
	return x.Svg(svgArgs...)
}
