package lucide

import (
	html "github.com/plainkit/html"
)

// PinOff creates a Pin Off Lucide icon.
func PinOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pin-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 17v5"))),
		html.Child(html.SvgPath(html.AD("M15 9.34V7a1 1 0 0 1 1-1 2 2 0 0 0 0-4H7.89"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M9 9v1.76a2 2 0 0 1-1.11 1.79l-1.78.9A2 2 0 0 0 5 15.24V16a1 1 0 0 0 1 1h11"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
