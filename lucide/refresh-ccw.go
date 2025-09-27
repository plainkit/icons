package lucide

import (
	html "github.com/plainkit/html"
)

// RefreshCcw creates a Refresh Ccw Lucide icon.
func RefreshCcw(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-refresh-ccw", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"))),
		html.Child(html.SvgPath(html.AD("M3 3v5h5"))),
		html.Child(html.SvgPath(html.AD("M3 12a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16"))),
		html.Child(html.SvgPath(html.AD("M16 16h5v5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
