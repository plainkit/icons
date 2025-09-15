package lucide

import x "github.com/plainkit/blox"

// BeerOff creates a Beer Off Lucide icon.
func BeerOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-beer-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 13v5"))),
		x.Child(x.Path(x.D("M17 11.47V8"))),
		x.Child(x.Path(x.D("M17 11h1a3 3 0 0 1 2.745 4.211"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M5 8v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-3"))),
		x.Child(x.Path(x.D("M7.536 7.535C6.766 7.649 6.154 8 5.5 8a2.5 2.5 0 0 1-1.768-4.268"))),
		x.Child(x.Path(x.D("M8.727 3.204C9.306 2.767 9.885 2 11 2c1.56 0 2 1.5 3 1.5s1.72-.5 2.5-.5a1 1 0 1 1 0 5c-.78 0-1.5-.5-2.5-.5a3.149 3.149 0 0 0-.842.12"))),
		x.Child(x.Path(x.D("M9 14.6V18"))),
	)
	return x.Svg(svgArgs...)
}
