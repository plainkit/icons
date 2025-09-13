package lucide

import x "github.com/bloxui/blox"

// BookPlus creates a Book Plus Lucide icon.
func BookPlus(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-book-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 7v6"))),
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Path(x.D("M9 10h6"))),
	)
	return x.Svg(svgArgs...)
}
