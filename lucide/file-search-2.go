package lucide

import (
	html "github.com/plainkit/html"
)

// FileSearch2 creates a File Search 2 Lucide icon.
func FileSearch2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-search-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgCircle(html.ACx("11.5"), html.ACy("14.5"), html.AR("2.5")),
		html.SvgPath(html.AD("M13.3 16.3 15 18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
