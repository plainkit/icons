package lucide

import x "github.com/plainkit/blox"

// BookCheck creates a Book Check Lucide icon.
func BookCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Path(x.D("m9 9.5 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
