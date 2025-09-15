package lucide

import x "github.com/plainkit/html"

// Trello creates a Trello Lucide icon.
func Trello(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-trello", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Rect(x.RectWidth("3"), x.RectHeight("9"), x.X("7"), x.Y("7"))),
		x.Child(x.Rect(x.RectWidth("3"), x.RectHeight("5"), x.X("14"), x.Y("7"))),
	)
	return x.Svg(svgArgs...)
}
