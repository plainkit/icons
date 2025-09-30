package lucide

import (
	html "github.com/plainkit/html"
)

// Pocket creates a Pocket Lucide icon.
func Pocket(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pocket", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20 3a2 2 0 0 1 2 2v6a1 1 0 0 1-20 0V5a2 2 0 0 1 2-2z")),
		html.SvgPath(html.AD("m8 10 4 4 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
