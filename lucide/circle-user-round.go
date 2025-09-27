package lucide

import (
	html "github.com/plainkit/html"
)

// CircleUserRound creates a Circle User Round Lucide icon.
func CircleUserRound(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-user-round", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 20a6 6 0 0 0-12 0"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("4"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
