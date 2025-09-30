package lucide

import (
	html "github.com/plainkit/html"
)

// InspectionPanel creates a Inspection Panel Lucide icon.
func InspectionPanel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-inspection-panel", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M7 7h.01")),
		html.SvgPath(html.AD("M17 7h.01")),
		html.SvgPath(html.AD("M7 17h.01")),
		html.SvgPath(html.AD("M17 17h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
