package lucide

import (
	html "github.com/plainkit/html"
)

// BedSingle creates a Bed Single Lucide icon.
func BedSingle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bed-single", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 20v-8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8"))),
		html.Child(html.SvgPath(html.AD("M5 10V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4"))),
		html.Child(html.SvgPath(html.AD("M3 18h18"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
