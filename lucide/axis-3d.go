package lucide

import (
	html "github.com/plainkit/html"
)

// Axis3d creates a Axis 3d Lucide icon.
func Axis3d(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-axis-3d", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13.5 10.5 15 9")),
		html.SvgPath(html.AD("M4 4v15a1 1 0 0 0 1 1h15")),
		html.SvgPath(html.AD("M4.293 19.707 6 18")),
		html.SvgPath(html.AD("m9 15 1.5-1.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
