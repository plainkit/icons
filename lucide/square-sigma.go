package lucide

import x "github.com/plainkit/blox"

// SquareSigma creates a Square Sigma Lucide icon.
func SquareSigma(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-sigma", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M16 8.9V7H8l4 5-4 5h8v-1.9"))),
	)
	return x.Svg(svgArgs...)
}
