package lucide

import (
	html "github.com/plainkit/html"
)

// UserRoundCheck creates a User Round Check Lucide icon.
func UserRoundCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-round-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 21a8 8 0 0 1 13.292-6")),
		html.SvgCircle(html.ACx("10"), html.ACy("8"), html.AR("5")),
		html.SvgPath(html.AD("m16 19 2 2 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
