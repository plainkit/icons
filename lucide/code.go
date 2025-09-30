package lucide

import (
	html "github.com/plainkit/html"
)

// Code creates a Code Lucide icon.
func Code(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-code", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m16 18 6-6-6-6")),
		html.SvgPath(html.AD("m8 6-6 6 6 6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
