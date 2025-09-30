package lucide

import (
	html "github.com/plainkit/html"
)

// Captions creates a Captions Lucide icon.
func Captions(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-captions", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("14"), html.AX("3"), html.AY("5"), html.ARx("2"), html.ARy("2")),
		html.SvgPath(html.AD("M7 15h4M15 15h2M7 11h2M13 11h4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
