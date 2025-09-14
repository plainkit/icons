package lucide

import x "github.com/bloxui/blox"

// Expand creates a Expand Lucide icon.
func Expand(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-expand", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 15 6 6"))),
		x.Child(x.Path(x.D("m15 9 6-6"))),
		x.Child(x.Path(x.D("M21 16v5h-5"))),
		x.Child(x.Path(x.D("M21 8V3h-5"))),
		x.Child(x.Path(x.D("M3 16v5h5"))),
		x.Child(x.Path(x.D("m3 21 6-6"))),
		x.Child(x.Path(x.D("M3 8V3h5"))),
		x.Child(x.Path(x.D("M9 9 3 3"))),
	)
	return x.Svg(svgArgs...)
}
