package lucide

import (
	html "github.com/plainkit/html"
)

// Shredder creates a Shredder Lucide icon.
func Shredder(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shredder", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 22v-5")),
		html.SvgPath(html.AD("M14 19v-2")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M18 20v-3")),
		html.SvgPath(html.AD("M2 13h20")),
		html.SvgPath(html.AD("M20 13V7l-5-5H6a2 2 0 0 0-2 2v9")),
		html.SvgPath(html.AD("M6 20v-3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
