package lucide

import x "github.com/bloxui/blox"

// VectorSquare creates a Vector Square Lucide icon.
func VectorSquare(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-vector-square", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19.5 7a24 24 0 0 1 0 10"))),
		x.Child(x.Path(x.D("M4.5 7a24 24 0 0 0 0 10"))),
		x.Child(x.Path(x.D("M7 19.5a24 24 0 0 0 10 0"))),
		x.Child(x.Path(x.D("M7 4.5a24 24 0 0 1 10 0"))),
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("5"), x.X("17"), x.Y("17"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("5"), x.X("17"), x.Y("2"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("5"), x.X("2"), x.Y("17"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("5"), x.X("2"), x.Y("2"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
