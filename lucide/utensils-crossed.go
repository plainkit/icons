package lucide

import x "github.com/bloxui/blox"

// UtensilsCrossed creates a Utensils Crossed Lucide icon.
func UtensilsCrossed(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-utensils-crossed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 2-2.3 2.3a3 3 0 0 0 0 4.2l1.8 1.8a3 3 0 0 0 4.2 0L22 8"))),
		x.Child(x.Path(x.D("M15 15 3.3 3.3a4.2 4.2 0 0 0 0 6l7.3 7.3c.7.7 2 .7 2.8 0L15 15Zm0 0 7 7"))),
		x.Child(x.Path(x.D("m2.1 21.8 6.4-6.3"))),
		x.Child(x.Path(x.D("m19 5-7 7"))),
	)
	return x.Svg(svgArgs...)
}
