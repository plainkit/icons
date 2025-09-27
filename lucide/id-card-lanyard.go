package lucide

import (
	html "github.com/plainkit/html"
)

// IdCardLanyard creates a Id Card Lanyard Lucide icon.
func IdCardLanyard(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-id-card-lanyard", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13.5 8h-3"))),
		html.Child(html.SvgPath(html.AD("m15 2-1 2h3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h3"))),
		html.Child(html.SvgPath(html.AD("M16.899 22A5 5 0 0 0 7.1 22"))),
		html.Child(html.SvgPath(html.AD("m9 2 3 6"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("15"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
