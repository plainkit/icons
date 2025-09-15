package lucide

import x "github.com/plainkit/blox"

// SquareParkingOff creates a Square Parking Off Lucide icon.
func SquareParkingOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-parking-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.6 3.6A2 2 0 0 1 5 3h14a2 2 0 0 1 2 2v14a2 2 0 0 1-.59 1.41"))),
		x.Child(x.Path(x.D("M3 8.7V19a2 2 0 0 0 2 2h10.3"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M13 13a3 3 0 1 0 0-6H9v2"))),
		x.Child(x.Path(x.D("M9 17v-2.3"))),
	)
	return x.Svg(svgArgs...)
}
