package lucide

import x "github.com/plainkit/html"

// SaudiRiyal creates a Saudi Riyal Lucide icon.
func SaudiRiyal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-saudi-riyal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m20 19.5-5.5 1.2"))),
		x.Child(x.Path(x.D("M14.5 4v11.22a1 1 0 0 0 1.242.97L20 15.2"))),
		x.Child(x.Path(x.D("m2.978 19.351 5.549-1.363A2 2 0 0 0 10 16V2"))),
		x.Child(x.Path(x.D("M20 10 4 13.5"))),
	)
	return x.Svg(svgArgs...)
}
