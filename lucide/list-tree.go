package lucide

import x "github.com/plainkit/blox"

// ListTree creates a List Tree Lucide icon.
func ListTree(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-tree", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 5h13"))),
		x.Child(x.Path(x.D("M13 12h8"))),
		x.Child(x.Path(x.D("M13 19h8"))),
		x.Child(x.Path(x.D("M3 10a2 2 0 0 0 2 2h3"))),
		x.Child(x.Path(x.D("M3 5v12a2 2 0 0 0 2 2h3"))),
	)
	return x.Svg(svgArgs...)
}
