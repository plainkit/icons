package lucide

import x "github.com/plainkit/blox"

// WifiLow creates a Wifi Low Lucide icon.
func WifiLow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wifi-low", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 20h.01"))),
		x.Child(x.Path(x.D("M8.5 16.429a5 5 0 0 1 7 0"))),
	)
	return x.Svg(svgArgs...)
}
