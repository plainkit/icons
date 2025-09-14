package lucide

import x "github.com/bloxui/blox"

// ClockAlert creates a Clock Alert Lucide icon.
func ClockAlert(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock-alert", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6l4 2"))),
		x.Child(x.Path(x.D("M20 12v5"))),
		x.Child(x.Path(x.D("M20 21h.01"))),
		x.Child(x.Path(x.D("M21.25 8.2A10 10 0 1 0 16 21.16"))),
	)
	return x.Svg(svgArgs...)
}
