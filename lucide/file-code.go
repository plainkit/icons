package lucide

import (
	html "github.com/plainkit/html"
)

// FileCode creates a File Code Lucide icon.
func FileCode(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-code", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 12.5 8 15l2 2.5"))),
		html.Child(html.SvgPath(html.AD("m14 12.5 2 2.5-2 2.5"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
