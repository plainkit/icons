package lucide

import (
	html "github.com/plainkit/html"
)

// FileOutput creates a File Output Lucide icon.
func FileOutput(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-output", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M4 7V4a2 2 0 0 1 2-2 2 2 0 0 0-2 2")),
		html.SvgPath(html.AD("M4.063 20.999a2 2 0 0 0 2 1L18 22a2 2 0 0 0 2-2V7l-5-5H6")),
		html.SvgPath(html.AD("m5 11-3 3")),
		html.SvgPath(html.AD("m5 17-3-3h10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
