package lucide

import x "github.com/bloxui/blox"

// Tent creates a Tent Lucide icon.
func Tent(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-tent", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.5 21 14 3"))),
		x.Child(x.Path(x.D("M20.5 21 10 3"))),
		x.Child(x.Path(x.D("M15.5 21 12 15l-3.5 6"))),
		x.Child(x.Path(x.D("M2 21h20"))),
	)
	return x.Svg(svgArgs...)
}
