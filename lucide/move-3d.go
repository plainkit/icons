package lucide

import (
	html "github.com/plainkit/html"
)

// Move3d creates a Move 3d Lucide icon.
func Move3d(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-3d", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 3v16h16"))),
		html.Child(html.SvgPath(html.AD("m5 19 6-6"))),
		html.Child(html.SvgPath(html.AD("m2 6 3-3 3 3"))),
		html.Child(html.SvgPath(html.AD("m18 16 3 3-3 3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
