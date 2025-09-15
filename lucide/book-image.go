package lucide

import x "github.com/plainkit/html"

// BookImage creates a Book Image Lucide icon.
func BookImage(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-image", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m20 13.7-2.1-2.1a2 2 0 0 0-2.8 0L9.7 17"))),
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("8"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
