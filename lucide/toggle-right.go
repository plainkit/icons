package lucide

import (
	html "github.com/plainkit/html"
)

// ToggleRight creates a Toggle Right Lucide icon.
func ToggleRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-toggle-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("15"), html.ACy("12"), html.AR("3"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("5"), html.ARx("7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
