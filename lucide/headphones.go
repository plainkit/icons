package lucide

import (
	html "github.com/plainkit/html"
)

// Headphones creates a Headphones Lucide icon.
func Headphones(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-headphones", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 14h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7a9 9 0 0 1 18 0v7a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
