package lucide

import x "github.com/plainkit/html"

// Captions creates a Captions Lucide icon.
func Captions(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-captions", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("14"), x.X("3"), x.Y("5"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M7 15h4M15 15h2M7 11h2M13 11h4"))),
	)
	return x.Svg(svgArgs...)
}
