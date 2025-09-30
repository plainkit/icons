package lucide

import (
	html "github.com/plainkit/html"
)

// Aperture creates a Aperture Lucide icon.
func Aperture(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-aperture", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("m14.31 8 5.74 9.94")),
		html.SvgPath(html.AD("M9.69 8h11.48")),
		html.SvgPath(html.AD("m7.38 12 5.74-9.94")),
		html.SvgPath(html.AD("M9.69 16 3.95 6.06")),
		html.SvgPath(html.AD("M14.31 16H2.83")),
		html.SvgPath(html.AD("m16.62 12-5.74 9.94")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
