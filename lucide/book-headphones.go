package lucide

import x "github.com/plainkit/html"

// BookHeadphones creates a Book Headphones Lucide icon.
func BookHeadphones(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-headphones", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Path(x.D("M8 12v-2a4 4 0 0 1 8 0v2"))),
		x.Child(x.Circle(x.Cx("15"), x.Cy("12"), x.R("1"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("12"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
