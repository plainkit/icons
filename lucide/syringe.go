package lucide

import (
	html "github.com/plainkit/html"
)

// Syringe creates a Syringe Lucide icon.
func Syringe(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-syringe", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m18 2 4 4")),
		html.SvgPath(html.AD("m17 7 3-3")),
		html.SvgPath(html.AD("M19 9 8.7 19.3c-1 1-2.5 1-3.4 0l-.6-.6c-1-1-1-2.5 0-3.4L15 5")),
		html.SvgPath(html.AD("m9 11 4 4")),
		html.SvgPath(html.AD("m5 19-3 3")),
		html.SvgPath(html.AD("m14 4 6 6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
