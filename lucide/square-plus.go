package lucide

import x "github.com/plainkit/html"

// SquarePlus creates a Square Plus Lucide icon.
func SquarePlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M8 12h8"))),
		x.Child(x.Path(x.D("M12 8v8"))),
	)
	return x.Svg(svgArgs...)
}
