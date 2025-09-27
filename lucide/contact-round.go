package lucide

import (
	html "github.com/plainkit/html"
)

// ContactRound creates a Contact Round Lucide icon.
func ContactRound(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-contact-round", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 2v2"))),
		html.Child(html.SvgPath(html.AD("M17.915 22a6 6 0 0 0-12 0"))),
		html.Child(html.SvgPath(html.AD("M8 2v2"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("4"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("4"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
