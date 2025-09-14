package lucide

import x "github.com/bloxui/blox"

// TabletSmartphone creates a Tablet Smartphone Lucide icon.
func TabletSmartphone(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tablet-smartphone", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("14"), x.X("3"), x.Y("8"), x.Rx("2"))),
		x.Child(x.Path(x.D("M5 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-2.4"))),
		x.Child(x.Path(x.D("M8 18h.01"))),
	)
	return x.Svg(svgArgs...)
}
