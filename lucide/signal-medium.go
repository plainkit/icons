package lucide

import x "github.com/bloxui/blox"

// SignalMedium creates a Signal Medium Lucide icon.
func SignalMedium(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-signal-medium", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 20h.01"))),
		x.Child(x.Path(x.D("M7 20v-4"))),
		x.Child(x.Path(x.D("M12 20v-8"))),
	)
	return x.Svg(svgArgs...)
}
