package lucide

import x "github.com/plainkit/html"

// BookLock creates a Book Lock Lucide icon.
func BookLock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-lock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 6V4a2 2 0 1 0-4 0v2"))),
		x.Child(x.Path(x.D("M20 15v6a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H10"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("5"), x.X("12"), x.Y("6"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
