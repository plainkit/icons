package lucide

import (
	html "github.com/plainkit/html"
)

// UserPen creates a User Pen Lucide icon.
func UserPen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-pen", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11.5 15H7a4 4 0 0 0-4 4v2"))),
		html.Child(html.SvgPath(html.AD("M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"))),
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("7"), html.AR("4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
