package lucide

import (
	html "github.com/plainkit/html"
)

// Maximize2 creates a Maximize 2 Lucide icon.
func Maximize2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-maximize-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 3h6v6")),
		html.SvgPath(html.AD("m21 3-7 7")),
		html.SvgPath(html.AD("m3 21 7-7")),
		html.SvgPath(html.AD("M9 21H3v-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
