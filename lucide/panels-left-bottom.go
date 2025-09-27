package lucide

import (
	html "github.com/plainkit/html"
)

// PanelsLeftBottom creates a Panels Left Bottom Lucide icon.
func PanelsLeftBottom(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panels-left-bottom", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M9 3v18"))),
		html.Child(html.SvgPath(html.AD("M9 15h12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
