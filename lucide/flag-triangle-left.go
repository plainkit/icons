package lucide

import (
	html "github.com/plainkit/html"
)

// FlagTriangleLeft creates a Flag Triangle Left Lucide icon.
func FlagTriangleLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flag-triangle-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18 22V2.8a.8.8 0 0 0-1.17-.71L5.45 7.78a.8.8 0 0 0 0 1.44L18 15.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
