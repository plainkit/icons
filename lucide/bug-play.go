package lucide

import (
	html "github.com/plainkit/html"
)

// BugPlay creates a Bug Play Lucide icon.
func BugPlay(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bug-play", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 19.655A6 6 0 0 1 6 14v-3a4 4 0 0 1 4-4h4a4 4 0 0 1 4 3.97")),
		html.SvgPath(html.AD("M14 15.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997a1 1 0 0 1-1.517-.86z")),
		html.SvgPath(html.AD("M14.12 3.88 16 2")),
		html.SvgPath(html.AD("M21 5a4 4 0 0 1-3.55 3.97")),
		html.SvgPath(html.AD("M3 21a4 4 0 0 1 3.81-4")),
		html.SvgPath(html.AD("M3 5a4 4 0 0 0 3.55 3.97")),
		html.SvgPath(html.AD("M6 13H2")),
		html.SvgPath(html.AD("m8 2 1.88 1.88")),
		html.SvgPath(html.AD("M9 7.13V6a3 3 0 1 1 6 0v1.13")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
