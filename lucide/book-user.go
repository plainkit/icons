package lucide

import (
	html "github.com/plainkit/html"
)

// BookUser creates a Book User Lucide icon.
func BookUser(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-user", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 13a3 3 0 1 0-6 0")),
		html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20")),
		html.SvgCircle(html.ACx("12"), html.ACy("8"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
