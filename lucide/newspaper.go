package lucide

import (
	html "github.com/plainkit/html"
)

// Newspaper creates a Newspaper Lucide icon.
func Newspaper(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-newspaper", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 18h-5"))),
		html.Child(html.SvgPath(html.AD("M18 14h-8"))),
		html.Child(html.SvgPath(html.AD("M4 22h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v16a2 2 0 0 1-4 0v-9a2 2 0 0 1 2-2h2"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("4"), html.AX("10"), html.AY("6"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
