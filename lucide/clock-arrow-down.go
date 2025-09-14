package lucide

import x "github.com/bloxui/blox"

// ClockArrowDown creates a Clock Arrow Down Lucide icon.
func ClockArrowDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock-arrow-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6l2 1"))),
		x.Child(x.Path(x.D("M12.337 21.994a10 10 0 1 1 9.588-8.767"))),
		x.Child(x.Path(x.D("m14 18 4 4 4-4"))),
		x.Child(x.Path(x.D("M18 14v8"))),
	)
	return x.Svg(svgArgs...)
}
