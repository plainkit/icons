package lucide

import (
	html "github.com/plainkit/html"
)

// Bomb creates a Bomb Lucide icon.
func Bomb(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bomb", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("11"), html.ACy("13"), html.AR("9")),
		html.SvgPath(html.AD("M14.35 4.65 16.3 2.7a2.41 2.41 0 0 1 3.4 0l1.6 1.6a2.4 2.4 0 0 1 0 3.4l-1.95 1.95")),
		html.SvgPath(html.AD("m22 2-1.5 1.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
