package lucide

import (
	html "github.com/plainkit/html"
)

// Tally1 creates a Tally 1 Lucide icon.
func Tally1(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tally-1", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 4v16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
