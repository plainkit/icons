package lucide

import (
	html "github.com/plainkit/html"
)

// Linkedin creates a Linkedin Lucide icon.
func Linkedin(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-linkedin", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"))),
		html.Child(html.SvgRect(html.AWidth("4"), html.AHeight("12"), html.AX("2"), html.AY("9"))),
		html.Child(html.SvgCircle(html.ACx("4"), html.ACy("4"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
