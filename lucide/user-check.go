package lucide

import (
	html "github.com/plainkit/html"
)

// UserCheck creates a User Check Lucide icon.
func UserCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m16 11 2 2 4-4")),
		html.SvgPath(html.AD("M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2")),
		html.SvgCircle(html.ACx("9"), html.ACy("7"), html.AR("4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
