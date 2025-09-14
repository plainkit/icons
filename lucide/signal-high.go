package lucide

import x "github.com/bloxui/blox"

// SignalHigh creates a Signal High Lucide icon.
func SignalHigh(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-signal-high", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 20h.01"))),
		x.Child(x.Path(x.D("M7 20v-4"))),
		x.Child(x.Path(x.D("M12 20v-8"))),
		x.Child(x.Path(x.D("M17 20V8"))),
	)
	return x.Svg(svgArgs...)
}
