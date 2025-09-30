package lucide

import (
	html "github.com/plainkit/html"
)

// WifiZero creates a Wifi Zero Lucide icon.
func WifiZero(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wifi-zero", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 20h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
