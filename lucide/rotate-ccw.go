package lucide

import (
	html "github.com/plainkit/html"
)

// RotateCcw creates a Rotate Ccw Lucide icon.
func RotateCcw(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rotate-ccw", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 12a9 9 0 1 0 9-9 9.75 9.75 0 0 0-6.74 2.74L3 8")),
		html.SvgPath(html.AD("M3 3v5h5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
