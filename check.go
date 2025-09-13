package lucide

import x "github.com/bloxui/blox"

// Check creates a Check Lucide icon.
func Check(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 6 9 17l-5-5"))),
	)
	return x.Svg(svgArgs...)
}
