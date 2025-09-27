package lucide

import (
	html "github.com/plainkit/html"
)

// PanelsTopLeft creates a Panels Top Left Lucide icon.
func PanelsTopLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panels-top-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M3 9h18"))),
		html.Child(html.SvgPath(html.AD("M9 21V9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
