package lucide

import (
	html "github.com/plainkit/html"
)

// Network creates a Network Lucide icon.
func Network(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-network", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("6"), html.AX("16"), html.AY("16"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("6"), html.AX("2"), html.AY("16"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("6"), html.AX("9"), html.AY("2"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("M5 16v-3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3"))),
		html.Child(html.SvgPath(html.AD("M12 12V8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
