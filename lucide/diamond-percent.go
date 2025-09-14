package lucide

import x "github.com/bloxui/blox"

// DiamondPercent creates a Diamond Percent Lucide icon.
func DiamondPercent(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-diamond-percent", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0Z"))),
		x.Child(x.Path(x.D("M9.2 9.2h.01"))),
		x.Child(x.Path(x.D("m14.5 9.5-5 5"))),
		x.Child(x.Path(x.D("M14.7 14.8h.01"))),
	)
	return x.Svg(svgArgs...)
}
