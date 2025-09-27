package lucide

import (
	html "github.com/plainkit/html"
)

// UserRoundPen creates a User Round Pen Lucide icon.
func UserRoundPen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-round-pen", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 21a8 8 0 0 1 10.821-7.487"))),
		html.Child(html.SvgPath(html.AD("M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"))),
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("8"), html.AR("5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
