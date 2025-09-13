package lucide

import x "github.com/bloxui/blox"

// BookA creates a Book A Lucide icon.
func BookA(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-book-a", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Path(x.D("m8 13 4-7 4 7"))),
		x.Child(x.Path(x.D("M9.1 11h5.7"))),
	)
	return x.Svg(svgArgs...)
}
