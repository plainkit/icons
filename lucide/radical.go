package lucide

import x "github.com/bloxui/blox"

// Radical creates a Radical Lucide icon.
func Radical(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-radical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 12h3.28a1 1 0 0 1 .948.684l2.298 7.934a.5.5 0 0 0 .96-.044L13.82 4.771A1 1 0 0 1 14.792 4H21"))),
	)
	return x.Svg(svgArgs...)
}
