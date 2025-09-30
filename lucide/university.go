package lucide

import (
	html "github.com/plainkit/html"
)

// University creates a University Lucide icon.
func University(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-university", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 21v-3a2 2 0 0 0-4 0v3")),
		html.SvgPath(html.AD("M18 12h.01")),
		html.SvgPath(html.AD("M18 16h.01")),
		html.SvgPath(html.AD("M22 7a1 1 0 0 0-1-1h-2a2 2 0 0 1-1.143-.359L13.143 2.36a2 2 0 0 0-2.286-.001L6.143 5.64A2 2 0 0 1 5 6H3a1 1 0 0 0-1 1v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2z")),
		html.SvgPath(html.AD("M6 12h.01")),
		html.SvgPath(html.AD("M6 16h.01")),
		html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
