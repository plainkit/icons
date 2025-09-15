package lucide

import x "github.com/plainkit/html"

// Rss creates a Rss Lucide icon.
func Rss(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rss", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 11a9 9 0 0 1 9 9"))),
		x.Child(x.Path(x.D("M4 4a16 16 0 0 1 16 16"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("19"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
