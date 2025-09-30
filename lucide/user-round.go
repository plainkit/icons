package lucide

import (
	html "github.com/plainkit/html"
)

// UserRound creates a User Round Lucide icon.
func UserRound(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-round", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("8"), html.AR("5")),
		html.SvgPath(html.AD("M20 21a8 8 0 0 0-16 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
