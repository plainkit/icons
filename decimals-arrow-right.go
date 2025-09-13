package lucide

import x "github.com/bloxui/blox"

// DecimalsArrowRight creates a Decimals Arrow Right Lucide icon.
func DecimalsArrowRight(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-decimals-arrow-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 18h10"))),
		x.Child(x.Path(x.D("m17 21 3-3-3-3"))),
		x.Child(x.Path(x.D("M3 11h.01"))),
		x.Child(x.Rect(x.X("15"), x.Y("3"), x.RectWidth("5"), x.RectHeight("8"), x.Rx("2.5"))),
		x.Child(x.Rect(x.X("6"), x.Y("3"), x.RectWidth("5"), x.RectHeight("8"), x.Rx("2.5"))),
	)
	return x.Svg(svgArgs...)
}
