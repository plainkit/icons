package lucide

import (
	html "github.com/plainkit/html"
)

// WalletMinimal creates a Wallet Minimal Lucide icon.
func WalletMinimal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wallet-minimal", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 14h.01")),
		html.SvgPath(html.AD("M7 7h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
