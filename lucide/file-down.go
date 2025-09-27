package lucide

import (
	html "github.com/plainkit/html"
)

// FileDown creates a File Down Lucide icon.
func FileDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-down", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M12 18v-6"))),
		html.Child(html.SvgPath(html.AD("m9 15 3 3 3-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
