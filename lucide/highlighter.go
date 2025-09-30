package lucide

import (
	html "github.com/plainkit/html"
)

// Highlighter creates a Highlighter Lucide icon.
func Highlighter(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-highlighter", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m9 11-6 6v3h9l3-3")),
		html.SvgPath(html.AD("m22 12-4.6 4.6a2 2 0 0 1-2.8 0l-5.2-5.2a2 2 0 0 1 0-2.8L14 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
