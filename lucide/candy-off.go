package lucide

import x "github.com/plainkit/html"

// CandyOff creates a Candy Off Lucide icon.
func CandyOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-candy-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 10v7.9"))),
		x.Child(x.Path(x.D("M11.802 6.145a5 5 0 0 1 6.053 6.053"))),
		x.Child(x.Path(x.D("M14 6.1v2.243"))),
		x.Child(x.Path(x.D("m15.5 15.571-.964.964a5 5 0 0 1-7.071 0 5 5 0 0 1 0-7.07l.964-.965"))),
		x.Child(x.Path(x.D("M16 7V3a1 1 0 0 1 1.707-.707 2.5 2.5 0 0 0 2.152.717 1 1 0 0 1 1.131 1.131 2.5 2.5 0 0 0 .717 2.152A1 1 0 0 1 21 8h-4"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M8 17v4a1 1 0 0 1-1.707.707 2.5 2.5 0 0 0-2.152-.717 1 1 0 0 1-1.131-1.131 2.5 2.5 0 0 0-.717-2.152A1 1 0 0 1 3 16h4"))),
	)
	return x.Svg(svgArgs...)
}
