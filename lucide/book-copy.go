package lucide

import (
	html "github.com/plainkit/html"
)

// BookCopy creates a Book Copy Lucide icon.
func BookCopy(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-copy", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 7a2 2 0 0 0-2 2v11")),
		html.SvgPath(html.AD("M5.803 18H5a2 2 0 0 0 0 4h9.5a.5.5 0 0 0 .5-.5V21")),
		html.SvgPath(html.AD("M9 15V4a2 2 0 0 1 2-2h9.5a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-.5.5H11a2 2 0 0 1 0-4h10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
