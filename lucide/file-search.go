package lucide

import (
	html "github.com/plainkit/html"
)

// FileSearch creates a File Search Lucide icon.
func FileSearch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-search", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M4.268 21a2 2 0 0 0 1.727 1H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3")),
		html.SvgPath(html.AD("m9 18-1.5-1.5")),
		html.SvgCircle(html.ACx("5"), html.ACy("14"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
