package lucide

import x "github.com/bloxui/blox"

// WalletCards creates a Wallet Cards Lucide icon.
func WalletCards(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wallet-cards", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 9a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M3 11h3c.8 0 1.6.3 2.1.9l1.1.9c1.6 1.6 4.1 1.6 5.7 0l1.1-.9c.5-.5 1.3-.9 2.1-.9H21"))),
	)
	return x.Svg(svgArgs...)
}
