package lucide

import (
	html "github.com/plainkit/html"
)

// Shirt creates a Shirt Lucide icon.
func Shirt(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shirt", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20.38 3.46 16 2a4 4 0 0 1-8 0L3.62 3.46a2 2 0 0 0-1.34 2.23l.58 3.47a1 1 0 0 0 .99.84H6v10c0 1.1.9 2 2 2h8a2 2 0 0 0 2-2V10h2.15a1 1 0 0 0 .99-.84l.58-3.47a2 2 0 0 0-1.34-2.23z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
