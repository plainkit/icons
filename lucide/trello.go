package lucide

import (
	html "github.com/plainkit/html"
)

// Trello creates a Trello Lucide icon.
func Trello(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-trello", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgRect(html.AWidth("3"), html.AHeight("9"), html.AX("7"), html.AY("7"))),
		html.Child(html.SvgRect(html.AWidth("3"), html.AHeight("5"), html.AX("14"), html.AY("7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
