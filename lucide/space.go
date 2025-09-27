package lucide

import (
	html "github.com/plainkit/html"
)

// Space creates a Space Lucide icon.
func Space(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-space", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 17v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
