package lucide

import (
	html "github.com/plainkit/html"
)

// Album creates a Album Lucide icon.
func Album(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-album", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"), html.ARy("2")),
		html.SvgPolyline(html.APoints("11 3 11 11 14 8 17 11 17 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
