package lucide

import (
	html "github.com/plainkit/html"
)

// UserSearch creates a User Search Lucide icon.
func UserSearch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-search", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("10"), html.ACy("7"), html.AR("4")),
		html.SvgPath(html.AD("M10.3 15H7a4 4 0 0 0-4 4v2")),
		html.SvgCircle(html.ACx("17"), html.ACy("17"), html.AR("3")),
		html.SvgPath(html.AD("m21 21-1.9-1.9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
