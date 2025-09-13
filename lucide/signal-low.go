package lucide

import x "github.com/bloxui/blox"

// SignalLow creates a Signal Low Lucide icon.
func SignalLow(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-signal-low", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 20h.01"))),
		x.Child(x.Path(x.D("M7 20v-4"))),
	)
	return x.Svg(svgArgs...)
}
