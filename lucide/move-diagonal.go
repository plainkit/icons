package lucide

import x "github.com/plainkit/html"

// MoveDiagonal creates a Move Diagonal Lucide icon.
func MoveDiagonal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-move-diagonal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 19H5v-6"))),
		x.Child(x.Path(x.D("M13 5h6v6"))),
		x.Child(x.Path(x.D("M19 5 5 19"))),
	)
	return x.Svg(svgArgs...)
}
