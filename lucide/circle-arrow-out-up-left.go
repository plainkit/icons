package lucide

import x "github.com/plainkit/blox"

// CircleArrowOutUpLeft creates a Circle Arrow Out Up Left Lucide icon.
func CircleArrowOutUpLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-arrow-out-up-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 8V2h6"))),
		x.Child(x.Path(x.D("m2 2 10 10"))),
		x.Child(x.Path(x.D("M12 2A10 10 0 1 1 2 12"))),
	)
	return x.Svg(svgArgs...)
}
