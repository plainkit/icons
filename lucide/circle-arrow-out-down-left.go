package lucide

import x "github.com/plainkit/blox"

// CircleArrowOutDownLeft creates a Circle Arrow Out Down Left Lucide icon.
func CircleArrowOutDownLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-arrow-out-down-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 12a10 10 0 1 1 10 10"))),
		x.Child(x.Path(x.D("m2 22 10-10"))),
		x.Child(x.Path(x.D("M8 22H2v-6"))),
	)
	return x.Svg(svgArgs...)
}
