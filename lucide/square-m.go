package lucide

import x "github.com/bloxui/blox"

// SquareM creates a Square M Lucide icon.
func SquareM(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-m", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 16V8.5a.5.5 0 0 1 .9-.3l2.7 3.599a.5.5 0 0 0 .8 0l2.7-3.6a.5.5 0 0 1 .9.3V16"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
