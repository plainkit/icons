package lucide

import x "github.com/plainkit/html"

// Newspaper creates a Newspaper Lucide icon.
func Newspaper(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-newspaper", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 18h-5"))),
		x.Child(x.Path(x.D("M18 14h-8"))),
		x.Child(x.Path(x.D("M4 22h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v16a2 2 0 0 1-4 0v-9a2 2 0 0 1 2-2h2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("4"), x.X("10"), x.Y("6"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
