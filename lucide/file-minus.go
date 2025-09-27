package lucide

import (
	html "github.com/plainkit/html"
)

// FileMinus creates a File Minus Lucide icon.
func FileMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-minus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M9 15h6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
