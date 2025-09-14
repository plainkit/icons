package lucide

import x "github.com/bloxui/blox"

// SquaresIntersect creates a Squares Intersect Lucide icon.
func SquaresIntersect(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-squares-intersect", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 22a2 2 0 0 1-2-2"))),
		x.Child(x.Path(x.D("M14 2a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M16 22h-2"))),
		x.Child(x.Path(x.D("M2 10V8"))),
		x.Child(x.Path(x.D("M2 4a2 2 0 0 1 2-2"))),
		x.Child(x.Path(x.D("M20 8a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M22 14v2"))),
		x.Child(x.Path(x.D("M22 20a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M4 16a2 2 0 0 1-2-2"))),
		x.Child(x.Path(x.D("M8 10a2 2 0 0 1 2-2h5a1 1 0 0 1 1 1v5a2 2 0 0 1-2 2H9a1 1 0 0 1-1-1z"))),
		x.Child(x.Path(x.D("M8 2h2"))),
	)
	return x.Svg(svgArgs...)
}
