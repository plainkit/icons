package lucide

import (
	html "github.com/plainkit/html"
)

// Orbit creates a Orbit Lucide icon.
func Orbit(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-orbit", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20.341 6.484A10 10 0 0 1 10.266 21.85")),
		html.SvgPath(html.AD("M3.659 17.516A10 10 0 0 1 13.74 2.152")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3")),
		html.SvgCircle(html.ACx("19"), html.ACy("5"), html.AR("2")),
		html.SvgCircle(html.ACx("5"), html.ACy("19"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
