package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronDown creates a Chevron Down Lucide icon.
func ChevronDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevron-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m6 9 6 6 6-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
