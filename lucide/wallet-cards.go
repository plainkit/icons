package lucide

import (
	html "github.com/plainkit/html"
)

// WalletCards creates a Wallet Cards Lucide icon.
func WalletCards(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wallet-cards", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M3 9a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2"))),
		html.Child(html.SvgPath(html.AD("M3 11h3c.8 0 1.6.3 2.1.9l1.1.9c1.6 1.6 4.1 1.6 5.7 0l1.1-.9c.5-.5 1.3-.9 2.1-.9H21"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
