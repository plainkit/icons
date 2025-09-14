package lucide

import x "github.com/bloxui/blox"

// SquarePlay creates a Square Play Lucide icon.
func SquarePlay(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-play", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M9 9.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997A1 1 0 0 1 9 14.996z"))),
	)
	return x.Svg(svgArgs...)
}
