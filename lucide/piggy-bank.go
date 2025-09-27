package lucide

import (
	html "github.com/plainkit/html"
)

// PiggyBank creates a Piggy Bank Lucide icon.
func PiggyBank(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-piggy-bank", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 17h3v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3a3.16 3.16 0 0 0 2-2h1a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-1a5 5 0 0 0-2-4V3a4 4 0 0 0-3.2 1.6l-.3.4H11a6 6 0 0 0-6 6v1a5 5 0 0 0 2 4v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1z"))),
		html.Child(html.SvgPath(html.AD("M16 10h.01"))),
		html.Child(html.SvgPath(html.AD("M2 8v1a2 2 0 0 0 2 2h1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
