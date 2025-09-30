package lucide

import (
	html "github.com/plainkit/html"
)

// RefreshCw creates a Refresh Cw Lucide icon.
func RefreshCw(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-refresh-cw", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 12a9 9 0 0 1 9-9 9.75 9.75 0 0 1 6.74 2.74L21 8")),
		html.SvgPath(html.AD("M21 3v5h-5")),
		html.SvgPath(html.AD("M21 12a9 9 0 0 1-9 9 9.75 9.75 0 0 1-6.74-2.74L3 16")),
		html.SvgPath(html.AD("M8 16H3v5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
