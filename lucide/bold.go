package lucide

import (
	html "github.com/plainkit/html"
)

// Bold creates a Bold Lucide icon.
func Bold(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bold", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 12h9a4 4 0 0 1 0 8H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h7a4 4 0 0 1 0 8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
