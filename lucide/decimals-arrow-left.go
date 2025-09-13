package lucide

import x "github.com/bloxui/blox"

// DecimalsArrowLeft creates a Decimals Arrow Left Lucide icon.
func DecimalsArrowLeft(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-decimals-arrow-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m13 21-3-3 3-3"))),
		x.Child(x.Path(x.D("M20 18H10"))),
		x.Child(x.Path(x.D("M3 11h.01"))),
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("8"), x.X("6"), x.Y("3"), x.Rx("2.5"))),
	)
	return x.Svg(svgArgs...)
}
