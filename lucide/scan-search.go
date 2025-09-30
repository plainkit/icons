package lucide

import (
	html "github.com/plainkit/html"
)

// ScanSearch creates a Scan Search Lucide icon.
func ScanSearch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scan-search", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 7V5a2 2 0 0 1 2-2h2")),
		html.SvgPath(html.AD("M17 3h2a2 2 0 0 1 2 2v2")),
		html.SvgPath(html.AD("M21 17v2a2 2 0 0 1-2 2h-2")),
		html.SvgPath(html.AD("M7 21H5a2 2 0 0 1-2-2v-2")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3")),
		html.SvgPath(html.AD("m16 16-1.9-1.9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
