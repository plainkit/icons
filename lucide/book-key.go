package lucide

import x "github.com/bloxui/blox"

// BookKey creates a Book Key Lucide icon.
func BookKey(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-book-key", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m19 3 1 1"))),
		x.Child(x.Path(x.D("m20 2-4.5 4.5"))),
		x.Child(x.Path(x.D("M20 7.898V21a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2h7.844"))),
		x.Child(x.Circle(x.Cx("14"), x.Cy("8"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
