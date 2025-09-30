package lucide

import (
	html "github.com/plainkit/html"
)

// Image creates a Image Lucide icon.
func Image(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-image", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"), html.ARy("2")),
		html.SvgCircle(html.ACx("9"), html.ACy("9"), html.AR("2")),
		html.SvgPath(html.AD("m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
