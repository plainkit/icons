package lucide

import x "github.com/bloxui/blox"

// SignalZero creates a Signal Zero Lucide icon.
func SignalZero(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-signal-zero", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 20h.01"))),
	)
	return x.Svg(svgArgs...)
}
