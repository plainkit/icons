package lucide

import x "github.com/plainkit/html"

// Binary creates a Binary Lucide icon.
func Binary(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-binary", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("6"), x.X("14"), x.Y("14"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("6"), x.X("6"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M6 20h4"))),
		x.Child(x.Path(x.D("M14 10h4"))),
		x.Child(x.Path(x.D("M6 14h2v6"))),
		x.Child(x.Path(x.D("M14 4h2v6"))),
	)
	return x.Svg(svgArgs...)
}
