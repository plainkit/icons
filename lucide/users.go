package lucide

import (
	html "github.com/plainkit/html"
)

// Users creates a Users Lucide icon.
func Users(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-users", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2")),
		html.SvgPath(html.AD("M16 3.128a4 4 0 0 1 0 7.744")),
		html.SvgPath(html.AD("M22 21v-2a4 4 0 0 0-3-3.87")),
		html.SvgCircle(html.ACx("9"), html.ACy("7"), html.AR("4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
