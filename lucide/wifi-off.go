package lucide

import (
	html "github.com/plainkit/html"
)

// WifiOff creates a Wifi Off Lucide icon.
func WifiOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wifi-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 20h.01")),
		html.SvgPath(html.AD("M8.5 16.429a5 5 0 0 1 7 0")),
		html.SvgPath(html.AD("M5 12.859a10 10 0 0 1 5.17-2.69")),
		html.SvgPath(html.AD("M19 12.859a10 10 0 0 0-2.007-1.523")),
		html.SvgPath(html.AD("M2 8.82a15 15 0 0 1 4.177-2.643")),
		html.SvgPath(html.AD("M22 8.82a15 15 0 0 0-11.288-3.764")),
		html.SvgPath(html.AD("m2 2 20 20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
