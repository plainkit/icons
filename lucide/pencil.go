package lucide

import x "github.com/bloxui/blox"

// Pencil creates a Pencil Lucide icon.
func Pencil(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pencil", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"))),
		x.Child(x.Path(x.D("m15 5 4 4"))),
	)
	return x.Svg(svgArgs...)
}
