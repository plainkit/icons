package lucide

import x "github.com/bloxui/blox"

// HeartOff creates a Heart Off Lucide icon.
func HeartOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-heart-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.5 4.893a5.5 5.5 0 0 1 1.091.931.56.56 0 0 0 .818 0A5.49 5.49 0 0 1 22 9.5c0 1.872-1.002 3.356-2.187 4.655"))),
		x.Child(x.Path(x.D("m16.967 16.967-3.459 3.346a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5a5.5 5.5 0 0 1 2.747-4.761"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
