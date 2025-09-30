package lucide

import (
	html "github.com/plainkit/html"
)

// Barrel creates a Barrel Lucide icon.
func Barrel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-barrel", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 3a41 41 0 0 0 0 18")),
		html.SvgPath(html.AD("M14 3a41 41 0 0 1 0 18")),
		html.SvgPath(html.AD("M17 3a2 2 0 0 1 1.68.92 15.25 15.25 0 0 1 0 16.16A2 2 0 0 1 17 21H7a2 2 0 0 1-1.68-.92 15.25 15.25 0 0 1 0-16.16A2 2 0 0 1 7 3z")),
		html.SvgPath(html.AD("M3.84 17h16.32")),
		html.SvgPath(html.AD("M3.84 7h16.32")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
