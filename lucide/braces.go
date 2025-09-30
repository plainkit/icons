package lucide

import (
	html "github.com/plainkit/html"
)

// Braces creates a Braces Lucide icon.
func Braces(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-braces", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 3H7a2 2 0 0 0-2 2v5a2 2 0 0 1-2 2 2 2 0 0 1 2 2v5c0 1.1.9 2 2 2h1")),
		html.SvgPath(html.AD("M16 21h1a2 2 0 0 0 2-2v-5c0-1.1.9-2 2-2a2 2 0 0 1-2-2V5a2 2 0 0 0-2-2h-1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
