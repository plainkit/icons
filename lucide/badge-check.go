package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeCheck creates a Badge Check Lucide icon.
func BadgeCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		html.Child(html.SvgPath(html.AD("m9 12 2 2 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
