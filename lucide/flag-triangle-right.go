package lucide

import (
	html "github.com/plainkit/html"
)

// FlagTriangleRight creates a Flag Triangle Right Lucide icon.
func FlagTriangleRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flag-triangle-right", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 22V2.8a.8.8 0 0 1 1.17-.71l11.38 5.69a.8.8 0 0 1 0 1.44L6 15.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
