package lucide

import x "github.com/bloxui/blox"

// MapPinOff creates a Map Pin Off Lucide icon.
func MapPinOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map-pin-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.75 7.09a3 3 0 0 1 2.16 2.16"))),
		x.Child(x.Path(x.D("M17.072 17.072c-1.634 2.17-3.527 3.912-4.471 4.727a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 1.432-4.568"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M8.475 2.818A8 8 0 0 1 20 10c0 1.183-.31 2.377-.81 3.533"))),
		x.Child(x.Path(x.D("M9.13 9.13a3 3 0 0 0 3.74 3.74"))),
	)
	return x.Svg(svgArgs...)
}
