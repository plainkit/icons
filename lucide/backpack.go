package lucide

import (
	html "github.com/plainkit/html"
)

// Backpack creates a Backpack Lucide icon.
func Backpack(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-backpack", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 10a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z")),
		html.SvgPath(html.AD("M8 10h8")),
		html.SvgPath(html.AD("M8 18h8")),
		html.SvgPath(html.AD("M8 22v-6a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v6")),
		html.SvgPath(html.AD("M9 6V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
