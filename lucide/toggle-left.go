package lucide

import (
	html "github.com/plainkit/html"
)

// ToggleLeft creates a Toggle Left Lucide icon.
func ToggleLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-toggle-left", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("9"), html.ACy("12"), html.AR("3")),
		html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("5"), html.ARx("7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
