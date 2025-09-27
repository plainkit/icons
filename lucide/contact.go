package lucide

import (
	html "github.com/plainkit/html"
)

// Contact creates a Contact Lucide icon.
func Contact(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-contact", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 2v2"))),
		html.Child(html.SvgPath(html.AD("M7 22v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2"))),
		html.Child(html.SvgPath(html.AD("M8 2v2"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("11"), html.AR("3"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("4"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
