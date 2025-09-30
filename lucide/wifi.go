package lucide

import (
	html "github.com/plainkit/html"
)

// Wifi creates a Wifi Lucide icon.
func Wifi(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wifi", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 20h.01")),
		html.SvgPath(html.AD("M2 8.82a15 15 0 0 1 20 0")),
		html.SvgPath(html.AD("M5 12.859a10 10 0 0 1 14 0")),
		html.SvgPath(html.AD("M8.5 16.429a5 5 0 0 1 7 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
