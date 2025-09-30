package lucide

import (
	html "github.com/plainkit/html"
)

// TabletSmartphone creates a Tablet Smartphone Lucide icon.
func TabletSmartphone(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tablet-smartphone", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("10"), html.AHeight("14"), html.AX("3"), html.AY("8"), html.ARx("2")),
		html.SvgPath(html.AD("M5 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-2.4")),
		html.SvgPath(html.AD("M8 18h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
