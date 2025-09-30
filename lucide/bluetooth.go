package lucide

import (
	html "github.com/plainkit/html"
)

// Bluetooth creates a Bluetooth Lucide icon.
func Bluetooth(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bluetooth", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m7 7 10 10-5 5V2l5 5L7 17")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
