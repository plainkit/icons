package lucide

import (
	html "github.com/plainkit/html"
)

// CirclePlay creates a Circle Play Lucide icon.
func CirclePlay(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-play", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M9 9.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997A1 1 0 0 1 9 14.996z"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
