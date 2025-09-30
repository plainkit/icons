package lucide

import (
	html "github.com/plainkit/html"
)

// Ungroup creates a Ungroup Lucide icon.
func Ungroup(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ungroup", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("8"), html.AHeight("6"), html.AX("5"), html.AY("4"), html.ARx("1")),
		html.SvgRect(html.AWidth("8"), html.AHeight("6"), html.AX("11"), html.AY("14"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
