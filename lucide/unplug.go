package lucide

import (
	html "github.com/plainkit/html"
)

// Unplug creates a Unplug Lucide icon.
func Unplug(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-unplug", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m19 5 3-3"))),
		html.Child(html.SvgPath(html.AD("m2 22 3-3"))),
		html.Child(html.SvgPath(html.AD("M6.3 20.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6-2.3 2.3a2.4 2.4 0 0 0 0 3.4Z"))),
		html.Child(html.SvgPath(html.AD("M7.5 13.5 10 11"))),
		html.Child(html.SvgPath(html.AD("M10.5 16.5 13 14"))),
		html.Child(html.SvgPath(html.AD("m12 6 6 6 2.3-2.3a2.4 2.4 0 0 0 0-3.4l-2.6-2.6a2.4 2.4 0 0 0-3.4 0Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
