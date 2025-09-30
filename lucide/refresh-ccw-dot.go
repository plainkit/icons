package lucide

import (
	html "github.com/plainkit/html"
)

// RefreshCcwDot creates a Refresh Ccw Dot Lucide icon.
func RefreshCcwDot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-refresh-ccw-dot", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8")),
		html.SvgPath(html.AD("M3 3v5h5")),
		html.SvgPath(html.AD("M3 12a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16")),
		html.SvgPath(html.AD("M16 16h5v5")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
