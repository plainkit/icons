package lucide

import (
	html "github.com/plainkit/html"
)

// LampWallDown creates a Lamp Wall Down Lucide icon.
func LampWallDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lamp-wall-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M19.929 18.629A1 1 0 0 1 19 20H9a1 1 0 0 1-.928-1.371l2-5A1 1 0 0 1 11 13h6a1 1 0 0 1 .928.629z")),
		html.SvgPath(html.AD("M6 3a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1z")),
		html.SvgPath(html.AD("M8 6h4a2 2 0 0 1 2 2v5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
