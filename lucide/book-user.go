package lucide

import x "github.com/bloxui/blox"

// BookUser creates a Book User Lucide icon.
func BookUser(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-user", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 13a3 3 0 1 0-6 0"))),
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("8"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
