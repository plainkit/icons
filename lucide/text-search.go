package lucide

import (
	html "github.com/plainkit/html"
)

// TextSearch creates a Text Search Lucide icon.
func TextSearch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-search", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 5H3")),
		html.SvgPath(html.AD("M10 12H3")),
		html.SvgPath(html.AD("M10 19H3")),
		html.SvgCircle(html.ACx("17"), html.ACy("15"), html.AR("3")),
		html.SvgPath(html.AD("m21 19-1.9-1.9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
