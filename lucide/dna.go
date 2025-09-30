package lucide

import (
	html "github.com/plainkit/html"
)

// Dna creates a Dna Lucide icon.
func Dna(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dna", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m10 16 1.5 1.5")),
		html.SvgPath(html.AD("m14 8-1.5-1.5")),
		html.SvgPath(html.AD("M15 2c-1.798 1.998-2.518 3.995-2.807 5.993")),
		html.SvgPath(html.AD("m16.5 10.5 1 1")),
		html.SvgPath(html.AD("m17 6-2.891-2.891")),
		html.SvgPath(html.AD("M2 15c6.667-6 13.333 0 20-6")),
		html.SvgPath(html.AD("m20 9 .891.891")),
		html.SvgPath(html.AD("M3.109 14.109 4 15")),
		html.SvgPath(html.AD("m6.5 12.5 1 1")),
		html.SvgPath(html.AD("m7 18 2.891 2.891")),
		html.SvgPath(html.AD("M9 22c1.798-1.998 2.518-3.995 2.807-5.993")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
