package lucide

import (
	html "github.com/plainkit/html"
)

// Section creates a Section Lucide icon.
func Section(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-section", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 5a4 3 0 0 0-8 0c0 4 8 3 8 7a4 3 0 0 1-8 0")),
		html.SvgPath(html.AD("M8 19a4 3 0 0 0 8 0c0-4-8-3-8-7a4 3 0 0 1 8 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
