package lucide

import (
	html "github.com/plainkit/html"
)

// UserRoundX creates a User Round X Lucide icon.
func UserRoundX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-round-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 21a8 8 0 0 1 11.873-7")),
		html.SvgCircle(html.ACx("10"), html.ACy("8"), html.AR("5")),
		html.SvgPath(html.AD("m17 17 5 5")),
		html.SvgPath(html.AD("m22 17-5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
