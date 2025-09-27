package lucide

import (
	html "github.com/plainkit/html"
)

// Usb creates a Usb Lucide icon.
func Usb(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-usb", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("7"), html.AR("1"))),
		html.Child(html.SvgCircle(html.ACx("4"), html.ACy("20"), html.AR("1"))),
		html.Child(html.SvgPath(html.AD("M4.7 19.3 19 5"))),
		html.Child(html.SvgPath(html.AD("m21 3-3 1 2 2Z"))),
		html.Child(html.SvgPath(html.AD("M9.26 7.68 5 12l2 5"))),
		html.Child(html.SvgPath(html.AD("m10 14 5 2 3.5-3.5"))),
		html.Child(html.SvgPath(html.AD("m18 12 1-1 1 1-1 1Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
