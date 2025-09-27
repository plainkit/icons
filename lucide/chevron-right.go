package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronRight creates a Chevron Right Lucide icon.
func ChevronRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevron-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m9 18 6-6-6-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
