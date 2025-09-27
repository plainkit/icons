package lucide

import (
	html "github.com/plainkit/html"
)

// SkipBack creates a Skip Back Lucide icon.
func SkipBack(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-skip-back", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M17.971 4.285A2 2 0 0 1 21 6v12a2 2 0 0 1-3.029 1.715l-9.997-5.998a2 2 0 0 1-.003-3.432z"))),
		html.Child(html.SvgPath(html.AD("M3 20V4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
