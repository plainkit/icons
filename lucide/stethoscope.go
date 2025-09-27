package lucide

import (
	html "github.com/plainkit/html"
)

// Stethoscope creates a Stethoscope Lucide icon.
func Stethoscope(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-stethoscope", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 2v2"))),
		html.Child(html.SvgPath(html.AD("M5 2v2"))),
		html.Child(html.SvgPath(html.AD("M5 3H4a2 2 0 0 0-2 2v4a6 6 0 0 0 12 0V5a2 2 0 0 0-2-2h-1"))),
		html.Child(html.SvgPath(html.AD("M8 15a6 6 0 0 0 12 0v-3"))),
		html.Child(html.SvgCircle(html.ACx("20"), html.ACy("10"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
