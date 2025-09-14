package lucide

import x "github.com/bloxui/blox"

// Cable creates a Cable Lucide icon.
func Cable(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cable", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 19a1 1 0 0 1-1-1v-2a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a1 1 0 0 1-1 1z"))),
		x.Child(x.Path(x.D("M17 21v-2"))),
		x.Child(x.Path(x.D("M19 14V6.5a1 1 0 0 0-7 0v11a1 1 0 0 1-7 0V10"))),
		x.Child(x.Path(x.D("M21 21v-2"))),
		x.Child(x.Path(x.D("M3 5V3"))),
		x.Child(x.Path(x.D("M4 10a2 2 0 0 1-2-2V6a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2z"))),
		x.Child(x.Path(x.D("M7 5V3"))),
	)
	return x.Svg(svgArgs...)
}
