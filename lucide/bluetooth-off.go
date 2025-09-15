package lucide

import x "github.com/plainkit/html"

// BluetoothOff creates a Bluetooth Off Lucide icon.
func BluetoothOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bluetooth-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m17 17-5 5V12l-5 5"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M14.5 9.5 17 7l-5-5v4.5"))),
	)
	return x.Svg(svgArgs...)
}
