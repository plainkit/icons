package lucide

import (
	html "github.com/plainkit/html"
)

// Link creates a Link Lucide icon.
func Link(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-link", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71")),
		html.SvgPath(html.AD("M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
