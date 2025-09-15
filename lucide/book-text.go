package lucide

import x "github.com/plainkit/blox"

// BookText creates a Book Text Lucide icon.
func BookText(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-text", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Path(x.D("M8 11h8"))),
		x.Child(x.Path(x.D("M8 7h6"))),
	)
	return x.Svg(svgArgs...)
}
