package lucide

import (
	html "github.com/plainkit/html"
)

// Citrus creates a Citrus Lucide icon.
func Citrus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-citrus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21.66 17.67a1.08 1.08 0 0 1-.04 1.6A12 12 0 0 1 4.73 2.38a1.1 1.1 0 0 1 1.61-.04z")),
		html.SvgPath(html.AD("M19.65 15.66A8 8 0 0 1 8.35 4.34")),
		html.SvgPath(html.AD("m14 10-5.5 5.5")),
		html.SvgPath(html.AD("M14 17.85V10H6.15")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
