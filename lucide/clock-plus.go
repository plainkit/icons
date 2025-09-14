package lucide

import x "github.com/bloxui/blox"

// ClockPlus creates a Clock Plus Lucide icon.
func ClockPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6l3.644 1.822"))),
		x.Child(x.Path(x.D("M16 19h6"))),
		x.Child(x.Path(x.D("M19 16v6"))),
		x.Child(x.Path(x.D("M21.92 13.267a10 10 0 1 0-8.653 8.653"))),
	)
	return x.Svg(svgArgs...)
}
