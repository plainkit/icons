package lucide

import x "github.com/bloxui/blox"

// Cigarette creates a Cigarette Lucide icon.
func Cigarette(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-cigarette", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 12H3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h14"))),
		x.Child(x.Path(x.D("M18 8c0-2.5-2-2.5-2-5"))),
		x.Child(x.Path(x.D("M21 16a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1"))),
		x.Child(x.Path(x.D("M22 8c0-2.5-2-2.5-2-5"))),
		x.Child(x.Path(x.D("M7 12v4"))),
	)
	return x.Svg(svgArgs...)
}
