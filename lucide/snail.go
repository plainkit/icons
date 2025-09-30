package lucide

import (
	html "github.com/plainkit/html"
)

// Snail creates a Snail Lucide icon.
func Snail(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-snail", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 13a6 6 0 1 0 12 0 4 4 0 1 0-8 0 2 2 0 0 0 4 0")),
		html.SvgCircle(html.ACx("10"), html.ACy("13"), html.AR("8")),
		html.SvgPath(html.AD("M2 21h12c4.4 0 8-3.6 8-8V7a2 2 0 1 0-4 0v6")),
		html.SvgPath(html.AD("M18 3 19.1 5.2")),
		html.SvgPath(html.AD("M22 3 20.9 5.2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
