package lucide

import (
	html "github.com/plainkit/html"
)

// Film creates a Film Lucide icon.
func Film(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-film", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M7 3v18"))),
		html.Child(html.SvgPath(html.AD("M3 7.5h4"))),
		html.Child(html.SvgPath(html.AD("M3 12h18"))),
		html.Child(html.SvgPath(html.AD("M3 16.5h4"))),
		html.Child(html.SvgPath(html.AD("M17 3v18"))),
		html.Child(html.SvgPath(html.AD("M17 7.5h4"))),
		html.Child(html.SvgPath(html.AD("M17 16.5h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
