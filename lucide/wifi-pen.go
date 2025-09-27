package lucide

import (
	html "github.com/plainkit/html"
)

// WifiPen creates a Wifi Pen Lucide icon.
func WifiPen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wifi-pen", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 8.82a15 15 0 0 1 20 0"))),
		html.Child(html.SvgPath(html.AD("M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"))),
		html.Child(html.SvgPath(html.AD("M5 12.859a10 10 0 0 1 10.5-2.222"))),
		html.Child(html.SvgPath(html.AD("M8.5 16.429a5 5 0 0 1 3-1.406"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
