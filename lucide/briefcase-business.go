package lucide

import (
	html "github.com/plainkit/html"
)

// BriefcaseBusiness creates a Briefcase Business Lucide icon.
func BriefcaseBusiness(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-briefcase-business", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 12h.01")),
		html.SvgPath(html.AD("M16 6V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2")),
		html.SvgPath(html.AD("M22 13a18.15 18.15 0 0 1-20 0")),
		html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("6"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
