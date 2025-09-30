package lucide

import (
	html "github.com/plainkit/html"
)

// Car creates a Car Lucide icon.
func Car(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-car", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M19 17h2c.6 0 1-.4 1-1v-3c0-.9-.7-1.7-1.5-1.9C18.7 10.6 16 10 16 10s-1.3-1.4-2.2-2.3c-.5-.4-1.1-.7-1.8-.7H5c-.6 0-1.1.4-1.4.9l-1.4 2.9A3.7 3.7 0 0 0 2 12v4c0 .6.4 1 1 1h2")),
		html.SvgCircle(html.ACx("7"), html.ACy("17"), html.AR("2")),
		html.SvgPath(html.AD("M9 17h6")),
		html.SvgCircle(html.ACx("17"), html.ACy("17"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
