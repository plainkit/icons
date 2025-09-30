package lucide

import (
	html "github.com/plainkit/html"
)

// Redo2 creates a Redo 2 Lucide icon.
func Redo2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-redo-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m15 14 5-5-5-5")),
		html.SvgPath(html.AD("M20 9H9.5A5.5 5.5 0 0 0 4 14.5A5.5 5.5 0 0 0 9.5 20H13")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
