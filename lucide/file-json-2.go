package lucide

import (
	html "github.com/plainkit/html"
)

// FileJson2 creates a File Json 2 Lucide icon.
func FileJson2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-json-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M4 12a1 1 0 0 0-1 1v1a1 1 0 0 1-1 1 1 1 0 0 1 1 1v1a1 1 0 0 0 1 1"))),
		html.Child(html.SvgPath(html.AD("M8 18a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1 1 1 0 0 1-1-1v-1a1 1 0 0 0-1-1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
