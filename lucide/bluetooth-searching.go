package lucide

import x "github.com/bloxui/blox"

// BluetoothSearching creates a Bluetooth Searching Lucide icon.
func BluetoothSearching(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bluetooth-searching", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m7 7 10 10-5 5V2l5 5L7 17"))),
		x.Child(x.Path(x.D("M20.83 14.83a4 4 0 0 0 0-5.66"))),
		x.Child(x.Path(x.D("M18 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
