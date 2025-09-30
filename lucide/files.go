package lucide

import (
	html "github.com/plainkit/html"
)

// Files creates a Files Lucide icon.
func Files(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-files", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2a2 2 0 0 1 1.414.586l4 4A2 2 0 0 1 21 8v7a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2z")),
		html.SvgPath(html.AD("M15 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M5 7a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h8a2 2 0 0 0 1.732-1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
