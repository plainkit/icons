package lucide

import x "github.com/plainkit/blox"

// Bluetooth creates a Bluetooth Lucide icon.
func Bluetooth(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bluetooth", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m7 7 10 10-5 5V2l5 5L7 17"))),
	)
	return x.Svg(svgArgs...)
}
