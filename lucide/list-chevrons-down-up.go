package lucide

import x "github.com/plainkit/blox"

// ListChevronsDownUp creates a List Chevrons Down Up Lucide icon.
func ListChevronsDownUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-chevrons-down-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 5h8"))),
		x.Child(x.Path(x.D("M3 12h8"))),
		x.Child(x.Path(x.D("M3 19h8"))),
		x.Child(x.Path(x.D("m15 5 3 3 3-3"))),
		x.Child(x.Path(x.D("m15 19 3-3 3 3"))),
	)
	return x.Svg(svgArgs...)
}
