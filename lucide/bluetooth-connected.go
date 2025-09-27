package lucide

import (
	html "github.com/plainkit/html"
)

// BluetoothConnected creates a Bluetooth Connected Lucide icon.
func BluetoothConnected(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bluetooth-connected", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m7 7 10 10-5 5V2l5 5L7 17"))),
		html.Child(html.SvgLine(html.AX1("18"), html.AX2("21"), html.AY1("12"), html.AY2("12"))),
		html.Child(html.SvgLine(html.AX1("3"), html.AX2("6"), html.AY1("12"), html.AY2("12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
