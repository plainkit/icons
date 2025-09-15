package lucide

import x "github.com/plainkit/html"

// Wallet creates a Wallet Lucide icon.
func Wallet(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wallet", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19 7V4a1 1 0 0 0-1-1H5a2 2 0 0 0 0 4h15a1 1 0 0 1 1 1v4h-3a2 2 0 0 0 0 4h3a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1"))),
		x.Child(x.Path(x.D("M3 5v14a2 2 0 0 0 2 2h15a1 1 0 0 0 1-1v-4"))),
	)
	return x.Svg(svgArgs...)
}
