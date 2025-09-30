package lucide

import (
	html "github.com/plainkit/html"
)

// Projector creates a Projector Lucide icon.
func Projector(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-projector", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 7 3 5")),
		html.SvgPath(html.AD("M9 6V3")),
		html.SvgPath(html.AD("m13 7 2-2")),
		html.SvgCircle(html.ACx("9"), html.ACy("13"), html.AR("3")),
		html.SvgPath(html.AD("M11.83 12H20a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h2.17")),
		html.SvgPath(html.AD("M16 16h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
