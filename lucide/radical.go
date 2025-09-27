package lucide

import (
	html "github.com/plainkit/html"
)

// Radical creates a Radical Lucide icon.
func Radical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-radical", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 12h3.28a1 1 0 0 1 .948.684l2.298 7.934a.5.5 0 0 0 .96-.044L13.82 4.771A1 1 0 0 1 14.792 4H21"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
