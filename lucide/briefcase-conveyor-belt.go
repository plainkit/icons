package lucide

import (
	html "github.com/plainkit/html"
)

// BriefcaseConveyorBelt creates a Briefcase Conveyor Belt Lucide icon.
func BriefcaseConveyorBelt(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-briefcase-conveyor-belt", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 20v2")),
		html.SvgPath(html.AD("M14 20v2")),
		html.SvgPath(html.AD("M18 20v2")),
		html.SvgPath(html.AD("M21 20H3")),
		html.SvgPath(html.AD("M6 20v2")),
		html.SvgPath(html.AD("M8 16V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v12")),
		html.SvgRect(html.AWidth("16"), html.AHeight("10"), html.AX("4"), html.AY("6"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
