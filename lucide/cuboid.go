package lucide

import (
	html "github.com/plainkit/html"
)

// Cuboid creates a Cuboid Lucide icon.
func Cuboid(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cuboid", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m21.12 6.4-6.05-4.06a2 2 0 0 0-2.17-.05L2.95 8.41a2 2 0 0 0-.95 1.7v5.82a2 2 0 0 0 .88 1.66l6.05 4.07a2 2 0 0 0 2.17.05l9.95-6.12a2 2 0 0 0 .95-1.7V8.06a2 2 0 0 0-.88-1.66Z"))),
		html.Child(html.SvgPath(html.AD("M10 22v-8L2.25 9.15"))),
		html.Child(html.SvgPath(html.AD("m10 14 11.77-6.87"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
