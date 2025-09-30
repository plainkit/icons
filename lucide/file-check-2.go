package lucide

import (
	html "github.com/plainkit/html"
)

// FileCheck2 creates a File Check 2 Lucide icon.
func FileCheck2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-check-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("m3 15 2 2 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
