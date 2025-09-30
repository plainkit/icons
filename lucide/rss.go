package lucide

import (
	html "github.com/plainkit/html"
)

// Rss creates a Rss Lucide icon.
func Rss(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rss", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 11a9 9 0 0 1 9 9")),
		html.SvgPath(html.AD("M4 4a16 16 0 0 1 16 16")),
		html.SvgCircle(html.ACx("5"), html.ACy("19"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
