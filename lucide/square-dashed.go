package lucide

import x "github.com/bloxui/blox"

// SquareDashed creates a Square Dashed Lucide icon.
func SquareDashed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 3a2 2 0 0 0-2 2"))),
		x.Child(x.Path(x.D("M19 3a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M21 19a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M5 21a2 2 0 0 1-2-2"))),
		x.Child(x.Path(x.D("M9 3h1"))),
		x.Child(x.Path(x.D("M9 21h1"))),
		x.Child(x.Path(x.D("M14 3h1"))),
		x.Child(x.Path(x.D("M14 21h1"))),
		x.Child(x.Path(x.D("M3 9v1"))),
		x.Child(x.Path(x.D("M21 9v1"))),
		x.Child(x.Path(x.D("M3 14v1"))),
		x.Child(x.Path(x.D("M21 14v1"))),
	)
	return x.Svg(svgArgs...)
}
