package lucide

import (
	html "github.com/plainkit/html"
)

// Scroll creates a Scroll Lucide icon.
func Scroll(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scroll", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M19 17V5a2 2 0 0 0-2-2H4"))),
		html.Child(html.SvgPath(html.AD("M8 21h12a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1H11a1 1 0 0 0-1 1v1a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v2a1 1 0 0 0 1 1h3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
