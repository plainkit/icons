package lucide

import x "github.com/plainkit/html"

// PinOff creates a Pin Off Lucide icon.
func PinOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pin-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 17v5"))),
		x.Child(x.Path(x.D("M15 9.34V7a1 1 0 0 1 1-1 2 2 0 0 0 0-4H7.89"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M9 9v1.76a2 2 0 0 1-1.11 1.79l-1.78.9A2 2 0 0 0 5 15.24V16a1 1 0 0 0 1 1h11"))),
	)
	return x.Svg(svgArgs...)
}
