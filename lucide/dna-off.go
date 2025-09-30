package lucide

import (
	html "github.com/plainkit/html"
)

// DnaOff creates a Dna Off Lucide icon.
func DnaOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dna-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2c-1.35 1.5-2.092 3-2.5 4.5L14 8")),
		html.SvgPath(html.AD("m17 6-2.891-2.891")),
		html.SvgPath(html.AD("M2 15c3.333-3 6.667-3 10-3")),
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("m20 9 .891.891")),
		html.SvgPath(html.AD("M22 9c-1.5 1.35-3 2.092-4.5 2.5l-1-1")),
		html.SvgPath(html.AD("M3.109 14.109 4 15")),
		html.SvgPath(html.AD("m6.5 12.5 1 1")),
		html.SvgPath(html.AD("m7 18 2.891 2.891")),
		html.SvgPath(html.AD("M9 22c1.35-1.5 2.092-3 2.5-4.5L10 16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
