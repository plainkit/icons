package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronLeft creates a Chevron Left Lucide icon.
func ChevronLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevron-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m15 18-6-6 6-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
