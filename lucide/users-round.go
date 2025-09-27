package lucide

import (
	html "github.com/plainkit/html"
)

// UsersRound creates a Users Round Lucide icon.
func UsersRound(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-users-round", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 21a8 8 0 0 0-16 0"))),
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("8"), html.AR("5"))),
		html.Child(html.SvgPath(html.AD("M22 20c0-3.37-2-6.5-4-8a5 5 0 0 0-.45-8.3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
