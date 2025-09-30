package lucide

import (
	html "github.com/plainkit/html"
)

// Briefcase creates a Briefcase Lucide icon.
func Briefcase(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-briefcase", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 20V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16")),
		html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("6"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
