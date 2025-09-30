package lucide

import (
	html "github.com/plainkit/html"
)

// BugOff creates a Bug Off Lucide icon.
func BugOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bug-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 20v-8")),
		html.SvgPath(html.AD("M14.12 3.88 16 2")),
		html.SvgPath(html.AD("M15 7.13V6a3 3 0 0 0-5.14-2.1L8 2")),
		html.SvgPath(html.AD("M18 12.34V11a4 4 0 0 0-4-4h-1.3")),
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M21 5a4 4 0 0 1-3.55 3.97")),
		html.SvgPath(html.AD("M22 13h-3.34")),
		html.SvgPath(html.AD("M3 21a4 4 0 0 1 3.81-4")),
		html.SvgPath(html.AD("M6 13H2")),
		html.SvgPath(html.AD("M7.7 7.7A4 4 0 0 0 6 11v3a6 6 0 0 0 11.13 3.13")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
