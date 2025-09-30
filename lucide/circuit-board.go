package lucide

import (
	html "github.com/plainkit/html"
)

// CircuitBoard creates a Circuit Board Lucide icon.
func CircuitBoard(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circuit-board", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M11 9h4a2 2 0 0 0 2-2V3")),
		html.SvgCircle(html.ACx("9"), html.ACy("9"), html.AR("2")),
		html.SvgPath(html.AD("M7 21v-4a2 2 0 0 1 2-2h4")),
		html.SvgCircle(html.ACx("15"), html.ACy("15"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
