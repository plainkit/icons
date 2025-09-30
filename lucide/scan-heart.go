package lucide

import (
	html "github.com/plainkit/html"
)

// ScanHeart creates a Scan Heart Lucide icon.
func ScanHeart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scan-heart", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 3h2a2 2 0 0 1 2 2v2")),
		html.SvgPath(html.AD("M21 17v2a2 2 0 0 1-2 2h-2")),
		html.SvgPath(html.AD("M3 7V5a2 2 0 0 1 2-2h2")),
		html.SvgPath(html.AD("M7 21H5a2 2 0 0 1-2-2v-2")),
		html.SvgPath(html.AD("M7.828 13.07A3 3 0 0 1 12 8.764a3 3 0 0 1 4.172 4.306l-3.447 3.62a1 1 0 0 1-1.449 0z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
