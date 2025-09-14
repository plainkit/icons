package lucide

import x "github.com/bloxui/blox"

// Armchair creates a Armchair Lucide icon.
func Armchair(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-armchair", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19 9V6a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v3"))),
		x.Child(x.Path(x.D("M3 16a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5a2 2 0 0 0-4 0v1.5a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5V11a2 2 0 0 0-4 0z"))),
		x.Child(x.Path(x.D("M5 18v2"))),
		x.Child(x.Path(x.D("M19 18v2"))),
	)
	return x.Svg(svgArgs...)
}
