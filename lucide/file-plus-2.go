package lucide

import (
	html "github.com/plainkit/html"
)

// FilePlus2 creates a File Plus 2 Lucide icon.
func FilePlus2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-plus-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M3 15h6"))),
		html.Child(html.SvgPath(html.AD("M6 12v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
