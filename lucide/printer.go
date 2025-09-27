package lucide

import (
	html "github.com/plainkit/html"
)

// Printer creates a Printer Lucide icon.
func Printer(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-printer", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"))),
		html.Child(html.SvgPath(html.AD("M6 9V3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v6"))),
		html.Child(html.SvgRect(html.AWidth("12"), html.AHeight("8"), html.AX("6"), html.AY("14"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
