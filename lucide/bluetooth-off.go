package lucide

import (
	html "github.com/plainkit/html"
)

// BluetoothOff creates a Bluetooth Off Lucide icon.
func BluetoothOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bluetooth-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m17 17-5 5V12l-5 5")),
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M14.5 9.5 17 7l-5-5v4.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
