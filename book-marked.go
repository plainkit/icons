package lucide

import x "github.com/bloxui/blox"

// BookMarked creates a Book Marked Lucide icon.
func BookMarked(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-book-marked", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 2v8l3-3 3 3V2"))),
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
	)
	return x.Svg(svgArgs...)
}