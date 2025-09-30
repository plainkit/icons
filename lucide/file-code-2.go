package lucide

import (
	html "github.com/plainkit/html"
)

// FileCode2 creates a File Code 2 Lucide icon.
func FileCode2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-code-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("m5 12-3 3 3 3")),
		html.SvgPath(html.AD("m9 18 3-3-3-3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
