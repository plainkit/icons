package lucide

import x "github.com/plainkit/blox"

// Landmark creates a Landmark Lucide icon.
func Landmark(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-landmark", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 18v-7"))),
		x.Child(x.Path(x.D("M11.12 2.198a2 2 0 0 1 1.76.006l7.866 3.847c.476.233.31.949-.22.949H3.474c-.53 0-.695-.716-.22-.949z"))),
		x.Child(x.Path(x.D("M14 18v-7"))),
		x.Child(x.Path(x.D("M18 18v-7"))),
		x.Child(x.Path(x.D("M3 22h18"))),
		x.Child(x.Path(x.D("M6 18v-7"))),
	)
	return x.Svg(svgArgs...)
}
