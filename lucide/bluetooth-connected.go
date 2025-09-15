package lucide

import x "github.com/plainkit/html"

// BluetoothConnected creates a Bluetooth Connected Lucide icon.
func BluetoothConnected(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bluetooth-connected", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m7 7 10 10-5 5V2l5 5L7 17"))),
		x.Child(x.Line(x.X1("18"), x.X2("21"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("3"), x.X2("6"), x.Y1("12"), x.Y2("12"))),
	)
	return x.Svg(svgArgs...)
}
