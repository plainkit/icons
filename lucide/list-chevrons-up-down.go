package lucide

import x "github.com/plainkit/html"

// ListChevronsUpDown creates a List Chevrons Up Down Lucide icon.
func ListChevronsUpDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-chevrons-up-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 5h8"))),
		x.Child(x.Path(x.D("M3 12h8"))),
		x.Child(x.Path(x.D("M3 19h8"))),
		x.Child(x.Path(x.D("m15 8 3-3 3 3"))),
		x.Child(x.Path(x.D("m15 16 3 3 3-3"))),
	)
	return x.Svg(svgArgs...)
}
