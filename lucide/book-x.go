package lucide

import x "github.com/plainkit/blox"

// BookX creates a Book X Lucide icon.
func BookX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14.5 7-5 5"))),
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Path(x.D("m9.5 7 5 5"))),
	)
	return x.Svg(svgArgs...)
}
