package lucide

import (
	html "github.com/plainkit/html"
)

// Drone creates a Drone Lucide icon.
func Drone(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-drone", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 10 7 7"))),
		html.Child(html.SvgPath(html.AD("m10 14-3 3"))),
		html.Child(html.SvgPath(html.AD("m14 10 3-3"))),
		html.Child(html.SvgPath(html.AD("m14 14 3 3"))),
		html.Child(html.SvgPath(html.AD("M14.205 4.139a4 4 0 1 1 5.439 5.863"))),
		html.Child(html.SvgPath(html.AD("M19.637 14a4 4 0 1 1-5.432 5.868"))),
		html.Child(html.SvgPath(html.AD("M4.367 10a4 4 0 1 1 5.438-5.862"))),
		html.Child(html.SvgPath(html.AD("M9.795 19.862a4 4 0 1 1-5.429-5.873"))),
		html.Child(html.SvgRect(html.AWidth("4"), html.AHeight("8"), html.AX("10"), html.AY("8"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
