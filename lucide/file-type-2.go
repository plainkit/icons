package lucide

import (
	html "github.com/plainkit/html"
)

// FileType2 creates a File Type 2 Lucide icon.
func FileType2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-type-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M2 13v-1h6v1"))),
		html.Child(html.SvgPath(html.AD("M5 12v6"))),
		html.Child(html.SvgPath(html.AD("M4 18h2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
