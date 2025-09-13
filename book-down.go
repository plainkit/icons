package lucide

import x "github.com/bloxui/blox"

// BookDown creates a Book Down Lucide icon.
func BookDown(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-book-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 13V7"))),
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Path(x.D("m9 10 3 3 3-3"))),
	)
	return x.Svg(svgArgs...)
}
