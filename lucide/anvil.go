package lucide

import (
	html "github.com/plainkit/html"
)

// Anvil creates a Anvil Lucide icon.
func Anvil(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-anvil", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7 10H6a4 4 0 0 1-4-4 1 1 0 0 1 1-1h4"))),
		html.Child(html.SvgPath(html.AD("M7 5a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1 7 7 0 0 1-7 7H8a1 1 0 0 1-1-1z"))),
		html.Child(html.SvgPath(html.AD("M9 12v5"))),
		html.Child(html.SvgPath(html.AD("M15 12v5"))),
		html.Child(html.SvgPath(html.AD("M5 20a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3 1 1 0 0 1-1 1H6a1 1 0 0 1-1-1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
