package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronUp creates a Chevron Up Lucide icon.
func ChevronUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevron-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m18 15-6-6-6 6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
