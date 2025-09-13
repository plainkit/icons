package lucide

import x "github.com/bloxui/blox"

// TrainTrack creates a Train Track Lucide icon.
func TrainTrack(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-train-track", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 17 17 2"))),
		x.Child(x.Path(x.D("m2 14 8 8"))),
		x.Child(x.Path(x.D("m5 11 8 8"))),
		x.Child(x.Path(x.D("m8 8 8 8"))),
		x.Child(x.Path(x.D("m11 5 8 8"))),
		x.Child(x.Path(x.D("m14 2 8 8"))),
		x.Child(x.Path(x.D("M7 22 22 7"))),
	)
	return x.Svg(svgArgs...)
}
