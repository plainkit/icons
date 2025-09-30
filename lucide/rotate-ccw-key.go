package lucide

import (
	html "github.com/plainkit/html"
)

// RotateCcwKey creates a Rotate Ccw Key Lucide icon.
func RotateCcwKey(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rotate-ccw-key", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m14.5 9.5 1 1")),
		html.SvgPath(html.AD("m15.5 8.5-4 4")),
		html.SvgPath(html.AD("M3 12a9 9 0 1 0 9-9 9.74 9.74 0 0 0-6.74 2.74L3 8")),
		html.SvgPath(html.AD("M3 3v5h5")),
		html.SvgCircle(html.ACx("10"), html.ACy("14"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
