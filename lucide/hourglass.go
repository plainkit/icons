package lucide

import (
	html "github.com/plainkit/html"
)

// Hourglass creates a Hourglass Lucide icon.
func Hourglass(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hourglass", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 22h14"))),
		html.Child(html.SvgPath(html.AD("M5 2h14"))),
		html.Child(html.SvgPath(html.AD("M17 22v-4.172a2 2 0 0 0-.586-1.414L12 12l-4.414 4.414A2 2 0 0 0 7 17.828V22"))),
		html.Child(html.SvgPath(html.AD("M7 2v4.172a2 2 0 0 0 .586 1.414L12 12l4.414-4.414A2 2 0 0 0 17 6.172V2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
