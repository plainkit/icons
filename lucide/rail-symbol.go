package lucide

import x "github.com/bloxui/blox"

// RailSymbol creates a Rail Symbol Lucide icon.
func RailSymbol(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-rail-symbol", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 15h14"))),
		x.Child(x.Path(x.D("M5 9h14"))),
		x.Child(x.Path(x.D("m14 20-5-5 6-6-5-5"))),
	)
	return x.Svg(svgArgs...)
}
