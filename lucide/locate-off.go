package lucide

import x "github.com/bloxui/blox"

// LocateOff creates a Locate Off Lucide icon.
func LocateOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-locate-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 19v3"))),
		x.Child(x.Path(x.D("M12 2v3"))),
		x.Child(x.Path(x.D("M18.89 13.24a7 7 0 0 0-8.13-8.13"))),
		x.Child(x.Path(x.D("M19 12h3"))),
		x.Child(x.Path(x.D("M2 12h3"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M7.05 7.05a7 7 0 0 0 9.9 9.9"))),
	)
	return x.Svg(svgArgs...)
}
