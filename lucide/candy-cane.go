package lucide

import (
	html "github.com/plainkit/html"
)

// CandyCane creates a Candy Cane Lucide icon.
func CandyCane(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-candy-cane", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5.7 21a2 2 0 0 1-3.5-2l8.6-14a6 6 0 0 1 10.4 6 2 2 0 1 1-3.464-2 2 2 0 1 0-3.464-2Z")),
		html.SvgPath(html.AD("M17.75 7 15 2.1")),
		html.SvgPath(html.AD("M10.9 4.8 13 9")),
		html.SvgPath(html.AD("m7.9 9.7 2 4.4")),
		html.SvgPath(html.AD("M4.9 14.7 7 18.9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
