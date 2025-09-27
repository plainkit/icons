package lucide

import (
	html "github.com/plainkit/html"
)

// Smartphone creates a Smartphone Lucide icon.
func Smartphone(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-smartphone", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("14"), html.AHeight("20"), html.AX("5"), html.AY("2"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgPath(html.AD("M12 18h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
