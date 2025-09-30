package lucide

import (
	html "github.com/plainkit/html"
)

// Pause creates a Pause Lucide icon.
func Pause(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pause", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("5"), html.AHeight("18"), html.AX("14"), html.AY("3"), html.ARx("1")),
		html.SvgRect(html.AWidth("5"), html.AHeight("18"), html.AX("5"), html.AY("3"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
