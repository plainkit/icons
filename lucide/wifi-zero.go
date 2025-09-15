package lucide

import x "github.com/plainkit/blox"

// WifiZero creates a Wifi Zero Lucide icon.
func WifiZero(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wifi-zero", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 20h.01"))),
	)
	return x.Svg(svgArgs...)
}
