package lucide

import (
	html "github.com/plainkit/html"
)

// BriefcaseMedical creates a Briefcase Medical Lucide icon.
func BriefcaseMedical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-briefcase-medical", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 11v4"))),
		html.Child(html.SvgPath(html.AD("M14 13h-4"))),
		html.Child(html.SvgPath(html.AD("M16 6V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2"))),
		html.Child(html.SvgPath(html.AD("M18 6v14"))),
		html.Child(html.SvgPath(html.AD("M6 6v14"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("6"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
