package lucide

import (
	html "github.com/plainkit/html"
)

// HardHat creates a Hard Hat Lucide icon.
func HardHat(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hard-hat", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 10V5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v5"))),
		html.Child(html.SvgPath(html.AD("M14 6a6 6 0 0 1 6 6v3"))),
		html.Child(html.SvgPath(html.AD("M4 15v-3a6 6 0 0 1 6-6"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("4"), html.AX("2"), html.AY("15"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
