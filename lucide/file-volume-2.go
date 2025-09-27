package lucide

import (
	html "github.com/plainkit/html"
)

// FileVolume2 creates a File Volume 2 Lucide icon.
func FileVolume2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-volume-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M8 15h.01"))),
		html.Child(html.SvgPath(html.AD("M11.5 13.5a2.5 2.5 0 0 1 0 3"))),
		html.Child(html.SvgPath(html.AD("M15 12a5 5 0 0 1 0 6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
