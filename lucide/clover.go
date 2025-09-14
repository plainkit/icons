package lucide

import x "github.com/bloxui/blox"

// Clover creates a Clover Lucide icon.
func Clover(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clover", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16.17 7.83 2 22"))),
		x.Child(x.Path(x.D("M4.02 12a2.827 2.827 0 1 1 3.81-4.17A2.827 2.827 0 1 1 12 4.02a2.827 2.827 0 1 1 4.17 3.81A2.827 2.827 0 1 1 19.98 12a2.827 2.827 0 1 1-3.81 4.17A2.827 2.827 0 1 1 12 19.98a2.827 2.827 0 1 1-4.17-3.81A1 1 0 1 1 4 12"))),
		x.Child(x.Path(x.D("m7.83 7.83 8.34 8.34"))),
	)
	return x.Svg(svgArgs...)
}
