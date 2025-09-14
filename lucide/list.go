package lucide

import x "github.com/bloxui/blox"

// List creates a List Lucide icon.
func List(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 5h.01"))),
		x.Child(x.Path(x.D("M3 12h.01"))),
		x.Child(x.Path(x.D("M3 19h.01"))),
		x.Child(x.Path(x.D("M8 5h13"))),
		x.Child(x.Path(x.D("M8 12h13"))),
		x.Child(x.Path(x.D("M8 19h13"))),
	)
	return x.Svg(svgArgs...)
}
