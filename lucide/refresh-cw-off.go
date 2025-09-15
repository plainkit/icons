package lucide

import x "github.com/plainkit/blox"

// RefreshCwOff creates a Refresh Cw Off Lucide icon.
func RefreshCwOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-refresh-cw-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 8L18.74 5.74A9.75 9.75 0 0 0 12 3C11 3 10.03 3.16 9.13 3.47"))),
		x.Child(x.Path(x.D("M8 16H3v5"))),
		x.Child(x.Path(x.D("M3 12C3 9.51 4 7.26 5.64 5.64"))),
		x.Child(x.Path(x.D("m3 16 2.26 2.26A9.75 9.75 0 0 0 12 21c2.49 0 4.74-1 6.36-2.64"))),
		x.Child(x.Path(x.D("M21 12c0 1-.16 1.97-.47 2.87"))),
		x.Child(x.Path(x.D("M21 3v5h-5"))),
		x.Child(x.Path(x.D("M22 22 2 2"))),
	)
	return x.Svg(svgArgs...)
}
