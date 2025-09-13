package lucide

import x "github.com/bloxui/blox"

// Gauge creates a Gauge Lucide icon.
func Gauge(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-gauge", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12 14 4-4"))),
		x.Child(x.Path(x.D("M3.34 19a10 10 0 1 1 17.32 0"))),
	)
	return x.Svg(svgArgs...)
}
