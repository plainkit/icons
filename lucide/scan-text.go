package lucide

import (
	html "github.com/plainkit/html"
)

// ScanText creates a Scan Text Lucide icon.
func ScanText(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scan-text", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 7V5a2 2 0 0 1 2-2h2")),
		html.SvgPath(html.AD("M17 3h2a2 2 0 0 1 2 2v2")),
		html.SvgPath(html.AD("M21 17v2a2 2 0 0 1-2 2h-2")),
		html.SvgPath(html.AD("M7 21H5a2 2 0 0 1-2-2v-2")),
		html.SvgPath(html.AD("M7 8h8")),
		html.SvgPath(html.AD("M7 12h10")),
		html.SvgPath(html.AD("M7 16h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
