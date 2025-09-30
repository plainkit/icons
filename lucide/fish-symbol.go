package lucide

import (
	html "github.com/plainkit/html"
)

// FishSymbol creates a Fish Symbol Lucide icon.
func FishSymbol(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-fish-symbol", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 16s9-15 20-4C11 23 2 8 2 8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
