package lucide

import (
	html "github.com/plainkit/html"
)

// FileInput creates a File Input Lucide icon.
func FileInput(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-input", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M2 15h10")),
		html.SvgPath(html.AD("m9 18 3-3-3-3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
