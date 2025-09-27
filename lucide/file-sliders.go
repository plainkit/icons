package lucide

import (
	html "github.com/plainkit/html"
)

// FileSliders creates a File Sliders Lucide icon.
func FileSliders(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-sliders", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M8 12h8"))),
		html.Child(html.SvgPath(html.AD("M10 11v2"))),
		html.Child(html.SvgPath(html.AD("M8 17h8"))),
		html.Child(html.SvgPath(html.AD("M14 16v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
