package lucide

import (
	html "github.com/plainkit/html"
)

// LifeBuoy creates a Life Buoy Lucide icon.
func LifeBuoy(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-life-buoy", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("m4.93 4.93 4.24 4.24"))),
		html.Child(html.SvgPath(html.AD("m14.83 9.17 4.24-4.24"))),
		html.Child(html.SvgPath(html.AD("m14.83 14.83 4.24 4.24"))),
		html.Child(html.SvgPath(html.AD("m9.17 14.83-4.24 4.24"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
