package lucide

import x "github.com/plainkit/html"

// CreditCard creates a Credit Card Lucide icon.
func CreditCard(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-credit-card", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("5"), x.Rx("2"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("10"), x.Y2("10"))),
	)
	return x.Svg(svgArgs...)
}
