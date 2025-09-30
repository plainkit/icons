package lucide

import (
	html "github.com/plainkit/html"
)

// FileX creates a File X Lucide icon.
func FileX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("m14.5 12.5-5 5")),
		html.SvgPath(html.AD("m9.5 12.5 5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
