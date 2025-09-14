package lucide

import x "github.com/bloxui/blox"

// WalletMinimal creates a Wallet Minimal Lucide icon.
func WalletMinimal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wallet-minimal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 14h.01"))),
		x.Child(x.Path(x.D("M7 7h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14"))),
	)
	return x.Svg(svgArgs...)
}
