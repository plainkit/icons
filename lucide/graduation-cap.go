package lucide

import (
	html "github.com/plainkit/html"
)

// GraduationCap creates a Graduation Cap Lucide icon.
func GraduationCap(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-graduation-cap", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21.42 10.922a1 1 0 0 0-.019-1.838L12.83 5.18a2 2 0 0 0-1.66 0L2.6 9.08a1 1 0 0 0 0 1.832l8.57 3.908a2 2 0 0 0 1.66 0z")),
		html.SvgPath(html.AD("M22 10v6")),
		html.SvgPath(html.AD("M6 12.5V16a6 3 0 0 0 12 0v-3.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
