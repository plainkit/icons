package lucide

import (
	html "github.com/plainkit/html"
)

// Vibrate creates a Vibrate Lucide icon.
func Vibrate(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-vibrate", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m2 8 2 2-2 2 2 2-2 2")),
		html.SvgPath(html.AD("m22 8-2 2 2 2-2 2 2 2")),
		html.SvgRect(html.AWidth("8"), html.AHeight("14"), html.AX("8"), html.AY("5"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
