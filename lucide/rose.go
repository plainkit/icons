package lucide

import (
	html "github.com/plainkit/html"
)

// Rose creates a Rose Lucide icon.
func Rose(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rose", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M17 10h-1a4 4 0 1 1 4-4v.534"))),
		html.Child(html.SvgPath(html.AD("M17 6h1a4 4 0 0 1 1.42 7.74l-2.29.87a6 6 0 0 1-5.339-10.68l2.069-1.31"))),
		html.Child(html.SvgPath(html.AD("M4.5 17c2.8-.5 4.4 0 5.5.8s1.8 2.2 2.3 3.7c-2 .4-3.5.4-4.8-.3-1.2-.6-2.3-1.9-3-4.2"))),
		html.Child(html.SvgPath(html.AD("M9.77 12C4 15 2 22 2 22"))),
		html.Child(html.SvgCircle(html.ACx("17"), html.ACy("8"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
