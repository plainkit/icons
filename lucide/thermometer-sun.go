package lucide

import x "github.com/bloxui/blox"

// ThermometerSun creates a Thermometer Sun Lucide icon.
func ThermometerSun(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-thermometer-sun", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 9a4 4 0 0 0-2 7.5"))),
		x.Child(x.Path(x.D("M12 3v2"))),
		x.Child(x.Path(x.D("m6.6 18.4-1.4 1.4"))),
		x.Child(x.Path(x.D("M20 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z"))),
		x.Child(x.Path(x.D("M4 13H2"))),
		x.Child(x.Path(x.D("M6.34 7.34 4.93 5.93"))),
	)
	return x.Svg(svgArgs...)
}
