package lucide

import (
	html "github.com/plainkit/html"
)

// ScanFace creates a Scan Face Lucide icon.
func ScanFace(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scan-face", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 7V5a2 2 0 0 1 2-2h2"))),
		html.Child(html.SvgPath(html.AD("M17 3h2a2 2 0 0 1 2 2v2"))),
		html.Child(html.SvgPath(html.AD("M21 17v2a2 2 0 0 1-2 2h-2"))),
		html.Child(html.SvgPath(html.AD("M7 21H5a2 2 0 0 1-2-2v-2"))),
		html.Child(html.SvgPath(html.AD("M8 14s1.5 2 4 2 4-2 4-2"))),
		html.Child(html.SvgPath(html.AD("M9 9h.01"))),
		html.Child(html.SvgPath(html.AD("M15 9h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
