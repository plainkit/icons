package lucide

import (
	html "github.com/plainkit/html"
)

// Beaker creates a Beaker Lucide icon.
func Beaker(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-beaker", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4.5 3h15")),
		html.SvgPath(html.AD("M6 3v16a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V3")),
		html.SvgPath(html.AD("M6 14h12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
