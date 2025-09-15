package lucide

import x "github.com/plainkit/blox"

// CarFront creates a Car Front Lucide icon.
func CarFront(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-car-front", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m21 8-2 2-1.5-3.7A2 2 0 0 0 15.646 5H8.4a2 2 0 0 0-1.903 1.257L5 10 3 8"))),
		x.Child(x.Path(x.D("M7 14h.01"))),
		x.Child(x.Path(x.D("M17 14h.01"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("8"), x.X("3"), x.Y("10"), x.Rx("2"))),
		x.Child(x.Path(x.D("M5 18v2"))),
		x.Child(x.Path(x.D("M19 18v2"))),
	)
	return x.Svg(svgArgs...)
}
