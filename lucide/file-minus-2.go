package lucide

import (
	html "github.com/plainkit/html"
)

// FileMinus2 creates a File Minus 2 Lucide icon.
func FileMinus2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-minus-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M3 15h6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
