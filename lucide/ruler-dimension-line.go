package lucide

import x "github.com/plainkit/html"

// RulerDimensionLine creates a Ruler Dimension Line Lucide icon.
func RulerDimensionLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ruler-dimension-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 15v-3.014"))),
		x.Child(x.Path(x.D("M16 15v-3.014"))),
		x.Child(x.Path(x.D("M20 6H4"))),
		x.Child(x.Path(x.D("M20 8V4"))),
		x.Child(x.Path(x.D("M4 8V4"))),
		x.Child(x.Path(x.D("M8 15v-3.014"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("7"), x.X("3"), x.Y("12"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
