package lucide

import x "github.com/plainkit/html"

// TentTree creates a Tent Tree Lucide icon.
func TentTree(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tent-tree", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("4"), x.Cy("4"), x.R("2"))),
		x.Child(x.Path(x.D("m14 5 3-3 3 3"))),
		x.Child(x.Path(x.D("m14 10 3-3 3 3"))),
		x.Child(x.Path(x.D("M17 14V2"))),
		x.Child(x.Path(x.D("M17 14H7l-5 8h20Z"))),
		x.Child(x.Path(x.D("M8 14v8"))),
		x.Child(x.Path(x.D("m9 14 5 8"))),
	)
	return x.Svg(svgArgs...)
}
