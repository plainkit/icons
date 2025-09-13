package lucide

import x "github.com/bloxui/blox"

// ChevronsUpDown creates a Chevrons Up Down Lucide icon.
func ChevronsUpDown(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chevrons-up-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m7 15 5 5 5-5"))),
		x.Child(x.Path(x.D("m7 9 5-5 5 5"))),
	)
	return x.Svg(svgArgs...)
}
