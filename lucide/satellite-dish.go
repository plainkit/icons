package lucide

import x "github.com/bloxui/blox"

// SatelliteDish creates a Satellite Dish Lucide icon.
func SatelliteDish(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-satellite-dish", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 10a7.31 7.31 0 0 0 10 10Z"))),
		x.Child(x.Path(x.D("m9 15 3-3"))),
		x.Child(x.Path(x.D("M17 13a6 6 0 0 0-6-6"))),
		x.Child(x.Path(x.D("M21 13A10 10 0 0 0 11 3"))),
	)
	return x.Svg(svgArgs...)
}
