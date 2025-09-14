package lucide

import x "github.com/bloxui/blox"

// BedSingle creates a Bed Single Lucide icon.
func BedSingle(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bed-single", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 20v-8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8"))),
		x.Child(x.Path(x.D("M5 10V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4"))),
		x.Child(x.Path(x.D("M3 18h18"))),
	)
	return x.Svg(svgArgs...)
}
