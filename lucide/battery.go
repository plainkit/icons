package lucide

import (
	html "github.com/plainkit/html"
)

// Battery creates a Battery Lucide icon.
func Battery(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-battery", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M 22 14 L 22 10")),
		html.SvgRect(html.AWidth("16"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
