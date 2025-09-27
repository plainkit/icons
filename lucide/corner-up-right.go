package lucide

import (
	html "github.com/plainkit/html"
)

// CornerUpRight creates a Corner Up Right Lucide icon.
func CornerUpRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-corner-up-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m15 14 5-5-5-5"))),
		html.Child(html.SvgPath(html.AD("M4 20v-7a4 4 0 0 1 4-4h12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
