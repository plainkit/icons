package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronLast creates a Chevron Last Lucide icon.
func ChevronLast(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevron-last", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m7 18 6-6-6-6")),
		html.SvgPath(html.AD("M17 6v12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
