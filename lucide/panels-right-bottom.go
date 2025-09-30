package lucide

import (
	html "github.com/plainkit/html"
)

// PanelsRightBottom creates a Panels Right Bottom Lucide icon.
func PanelsRightBottom(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panels-right-bottom", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M3 15h12")),
		html.SvgPath(html.AD("M15 3v18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
