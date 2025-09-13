package lucide

import x "github.com/bloxui/blox"

// EggOff creates a Egg Off Lucide icon.
func EggOff(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-egg-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M20 14.347V14c0-6-4-12-8-12-1.078 0-2.157.436-3.157 1.19"))),
		x.Child(x.Path(x.D("M6.206 6.21C4.871 8.4 4 11.2 4 14a8 8 0 0 0 14.568 4.568"))),
	)
	return x.Svg(svgArgs...)
}
