package lucide

import (
	html "github.com/plainkit/html"
)

// TextInitial creates a Text Initial Lucide icon.
func TextInitial(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-initial", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 5h6")),
		html.SvgPath(html.AD("M15 12h6")),
		html.SvgPath(html.AD("M3 19h18")),
		html.SvgPath(html.AD("m3 12 3.553-7.724a.5.5 0 0 1 .894 0L11 12")),
		html.SvgPath(html.AD("M3.92 10h6.16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
