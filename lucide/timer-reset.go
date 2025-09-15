package lucide

import x "github.com/plainkit/html"

// TimerReset creates a Timer Reset Lucide icon.
func TimerReset(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-timer-reset", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 2h4"))),
		x.Child(x.Path(x.D("M12 14v-4"))),
		x.Child(x.Path(x.D("M4 13a8 8 0 0 1 8-7 8 8 0 1 1-5.3 14L4 17.6"))),
		x.Child(x.Path(x.D("M9 17H4v5"))),
	)
	return x.Svg(svgArgs...)
}
