package lucide

import (
	html "github.com/plainkit/html"
)

// WifiLow creates a Wifi Low Lucide icon.
func WifiLow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wifi-low", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 20h.01")),
		html.SvgPath(html.AD("M8.5 16.429a5 5 0 0 1 7 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
