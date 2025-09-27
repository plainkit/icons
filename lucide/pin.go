package lucide

import (
	html "github.com/plainkit/html"
)

// Pin creates a Pin Lucide icon.
func Pin(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pin", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 17v5"))),
		html.Child(html.SvgPath(html.AD("M9 10.76a2 2 0 0 1-1.11 1.79l-1.78.9A2 2 0 0 0 5 15.24V16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-.76a2 2 0 0 0-1.11-1.79l-1.78-.9A2 2 0 0 1 15 10.76V7a1 1 0 0 1 1-1 2 2 0 0 0 0-4H8a2 2 0 0 0 0 4 1 1 0 0 1 1 1z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
