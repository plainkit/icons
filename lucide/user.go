package lucide

import (
	html "github.com/plainkit/html"
)

// User creates a User Lucide icon.
func User(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2")),
		html.SvgCircle(html.ACx("12"), html.ACy("7"), html.AR("4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
