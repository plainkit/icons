package lucide

import (
	html "github.com/plainkit/html"
)

// CreditCard creates a Credit Card Lucide icon.
func CreditCard(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-credit-card", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("5"), html.ARx("2"))),
		html.Child(html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("10"), html.AY2("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
