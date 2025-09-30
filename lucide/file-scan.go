package lucide

import (
	html "github.com/plainkit/html"
)

// FileScan creates a File Scan Lucide icon.
func FileScan(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-scan", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20 10V7l-5-5H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M16 14a2 2 0 0 0-2 2")),
		html.SvgPath(html.AD("M20 14a2 2 0 0 1 2 2")),
		html.SvgPath(html.AD("M20 22a2 2 0 0 0 2-2")),
		html.SvgPath(html.AD("M16 22a2 2 0 0 1-2-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
