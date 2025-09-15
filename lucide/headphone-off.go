package lucide

import x "github.com/plainkit/blox"

// HeadphoneOff creates a Headphone Off Lucide icon.
func HeadphoneOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-headphone-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 14h-1.343"))),
		x.Child(x.Path(x.D("M9.128 3.47A9 9 0 0 1 21 12v3.343"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M20.414 20.414A2 2 0 0 1 19 21h-1a2 2 0 0 1-2-2v-3"))),
		x.Child(x.Path(x.D("M3 14h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7a9 9 0 0 1 2.636-6.364"))),
	)
	return x.Svg(svgArgs...)
}
