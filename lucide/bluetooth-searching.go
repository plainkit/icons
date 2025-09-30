package lucide

import (
	html "github.com/plainkit/html"
)

// BluetoothSearching creates a Bluetooth Searching Lucide icon.
func BluetoothSearching(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bluetooth-searching", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m7 7 10 10-5 5V2l5 5L7 17")),
		html.SvgPath(html.AD("M20.83 14.83a4 4 0 0 0 0-5.66")),
		html.SvgPath(html.AD("M18 12h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
