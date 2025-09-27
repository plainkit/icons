package lucide

import (
	html "github.com/plainkit/html"
)

// AtSign creates a At Sign Lucide icon.
func AtSign(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-at-sign", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("4"))),
		html.Child(html.SvgPath(html.AD("M16 8v5a3 3 0 0 0 6 0v-1a10 10 0 1 0-4 8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
