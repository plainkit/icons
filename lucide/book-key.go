package lucide

import (
	html "github.com/plainkit/html"
)

// BookKey creates a Book Key Lucide icon.
func BookKey(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-key", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m19 3 1 1")),
		html.SvgPath(html.AD("m20 2-4.5 4.5")),
		html.SvgPath(html.AD("M20 7.898V21a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20")),
		html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2h7.844")),
		html.SvgCircle(html.ACx("14"), html.ACy("8"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
