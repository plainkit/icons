package lucide

import (
	html "github.com/plainkit/html"
)

// Mails creates a Mails Lucide icon.
func Mails(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mails", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 1-1.732")),
		html.SvgPath(html.AD("m22 5.5-6.419 4.179a2 2 0 0 1-2.162 0L7 5.5")),
		html.SvgRect(html.AWidth("15"), html.AHeight("12"), html.AX("7"), html.AY("3"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
