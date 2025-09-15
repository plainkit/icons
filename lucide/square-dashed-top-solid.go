package lucide

import x "github.com/plainkit/html"

// SquareDashedTopSolid creates a Square Dashed Top Solid Lucide icon.
func SquareDashedTopSolid(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-dashed-top-solid", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 21h1"))),
		x.Child(x.Path(x.D("M21 14v1"))),
		x.Child(x.Path(x.D("M21 19a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M21 9v1"))),
		x.Child(x.Path(x.D("M3 14v1"))),
		x.Child(x.Path(x.D("M3 5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M3 9v1"))),
		x.Child(x.Path(x.D("M5 21a2 2 0 0 1-2-2"))),
		x.Child(x.Path(x.D("M9 21h1"))),
	)
	return x.Svg(svgArgs...)
}
