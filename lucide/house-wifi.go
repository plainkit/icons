package lucide

import x "github.com/bloxui/blox"

// HouseWifi creates a House Wifi Lucide icon.
func HouseWifi(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-house-wifi", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9.5 13.866a4 4 0 0 1 5 .01"))),
		x.Child(x.Path(x.D("M12 17h.01"))),
		x.Child(x.Path(x.D("M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"))),
		x.Child(x.Path(x.D("M7 10.754a8 8 0 0 1 10 0"))),
	)
	return x.Svg(svgArgs...)
}
