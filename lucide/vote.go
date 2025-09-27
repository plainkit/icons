package lucide

import (
	html "github.com/plainkit/html"
)

// Vote creates a Vote Lucide icon.
func Vote(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-vote", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m9 12 2 2 4-4"))),
		html.Child(html.SvgPath(html.AD("M5 7c0-1.1.9-2 2-2h10a2 2 0 0 1 2 2v12H5V7Z"))),
		html.Child(html.SvgPath(html.AD("M22 19H2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
