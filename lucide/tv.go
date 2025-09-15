package lucide

import x "github.com/plainkit/html"

// Tv creates a Tv Lucide icon.
func Tv(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tv", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m17 2-5 5-5-5"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("15"), x.X("2"), x.Y("7"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
