package lucide

import (
	html "github.com/plainkit/html"
)

// FilePenLine creates a File Pen Line Lucide icon.
func FilePenLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-pen-line", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m18 5-2.414-2.414A2 2 0 0 0 14.172 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2"))),
		html.Child(html.SvgPath(html.AD("M21.378 12.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"))),
		html.Child(html.SvgPath(html.AD("M8 18h1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
