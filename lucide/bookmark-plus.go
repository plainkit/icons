package lucide

import x "github.com/bloxui/blox"

// BookmarkPlus creates a Bookmark Plus Lucide icon.
func BookmarkPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bookmark-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16z"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("7"), x.Y2("13"))),
		x.Child(x.Line(x.X1("15"), x.X2("9"), x.Y1("10"), x.Y2("10"))),
	)
	return x.Svg(svgArgs...)
}
