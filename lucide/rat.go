package lucide

import (
	html "github.com/plainkit/html"
)

// Rat creates a Rat Lucide icon.
func Rat(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rat", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13 22H4a2 2 0 0 1 0-4h12")),
		html.SvgPath(html.AD("M13.236 18a3 3 0 0 0-2.2-5")),
		html.SvgPath(html.AD("M16 9h.01")),
		html.SvgPath(html.AD("M16.82 3.94a3 3 0 1 1 3.237 4.868l1.815 2.587a1.5 1.5 0 0 1-1.5 2.1l-2.872-.453a3 3 0 0 0-3.5 3")),
		html.SvgPath(html.AD("M17 4.988a3 3 0 1 0-5.2 2.052A7 7 0 0 0 4 14.015 4 4 0 0 0 8 18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
