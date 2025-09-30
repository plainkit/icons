package lucide

import (
	html "github.com/plainkit/html"
)

// CircleArrowLeft creates a Circle Arrow Left Lucide icon.
func CircleArrowLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-arrow-left", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("m12 8-4 4 4 4")),
		html.SvgPath(html.AD("M16 12H8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
