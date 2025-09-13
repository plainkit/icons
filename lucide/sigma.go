package lucide

import x "github.com/bloxui/blox"

// Sigma creates a Sigma Lucide icon.
func Sigma(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-sigma", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 7V5a1 1 0 0 0-1-1H6.5a.5.5 0 0 0-.4.8l4.5 6a2 2 0 0 1 0 2.4l-4.5 6a.5.5 0 0 0 .4.8H17a1 1 0 0 0 1-1v-2"))),
	)
	return x.Svg(svgArgs...)
}
