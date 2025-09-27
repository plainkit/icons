package lucide

import (
	html "github.com/plainkit/html"
)

// UserRoundMinus creates a User Round Minus Lucide icon.
func UserRoundMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-round-minus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 21a8 8 0 0 1 13.292-6"))),
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("8"), html.AR("5"))),
		html.Child(html.SvgPath(html.AD("M22 19h-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
