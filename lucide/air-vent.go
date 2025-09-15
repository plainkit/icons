package lucide

import x "github.com/plainkit/blox"

// AirVent creates a Air Vent Lucide icon.
func AirVent(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-air-vent", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 17.5a2.5 2.5 0 1 1-4 2.03V12"))),
		x.Child(x.Path(x.D("M6 12H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M6 8h12"))),
		x.Child(x.Path(x.D("M6.6 15.572A2 2 0 1 0 10 17v-5"))),
	)
	return x.Svg(svgArgs...)
}
