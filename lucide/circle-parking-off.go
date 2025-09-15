package lucide

import x "github.com/plainkit/blox"

// CircleParkingOff creates a Circle Parking Off Lucide icon.
func CircleParkingOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-parking-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.656 7H13a3 3 0 0 1 2.984 3.307"))),
		x.Child(x.Path(x.D("M13 13H9"))),
		x.Child(x.Path(x.D("M19.071 19.071A1 1 0 0 1 4.93 4.93"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M8.357 2.687a10 10 0 0 1 12.956 12.956"))),
		x.Child(x.Path(x.D("M9 17V9"))),
	)
	return x.Svg(svgArgs...)
}
