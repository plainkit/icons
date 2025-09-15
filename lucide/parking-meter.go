package lucide

import x "github.com/plainkit/blox"

// ParkingMeter creates a Parking Meter Lucide icon.
func ParkingMeter(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-parking-meter", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 15h2"))),
		x.Child(x.Path(x.D("M12 12v3"))),
		x.Child(x.Path(x.D("M12 19v3"))),
		x.Child(x.Path(x.D("M15.282 19a1 1 0 0 0 .948-.68l2.37-6.988a7 7 0 1 0-13.2 0l2.37 6.988a1 1 0 0 0 .948.68z"))),
		x.Child(x.Path(x.D("M9 9a3 3 0 1 1 6 0"))),
	)
	return x.Svg(svgArgs...)
}
