package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronFirst creates a Chevron First Lucide icon.
func ChevronFirst(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevron-first", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m17 18-6-6 6-6"))),
		html.Child(html.SvgPath(html.AD("M7 6v12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
