package lucide

import (
	html "github.com/plainkit/html"
)

// Baseline creates a Baseline Lucide icon.
func Baseline(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-baseline", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 20h16")),
		html.SvgPath(html.AD("m6 16 6-12 6 12")),
		html.SvgPath(html.AD("M8 12h8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
