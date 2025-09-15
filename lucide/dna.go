package lucide

import x "github.com/plainkit/html"

// Dna creates a Dna Lucide icon.
func Dna(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-dna", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 16 1.5 1.5"))),
		x.Child(x.Path(x.D("m14 8-1.5-1.5"))),
		x.Child(x.Path(x.D("M15 2c-1.798 1.998-2.518 3.995-2.807 5.993"))),
		x.Child(x.Path(x.D("m16.5 10.5 1 1"))),
		x.Child(x.Path(x.D("m17 6-2.891-2.891"))),
		x.Child(x.Path(x.D("M2 15c6.667-6 13.333 0 20-6"))),
		x.Child(x.Path(x.D("m20 9 .891.891"))),
		x.Child(x.Path(x.D("M3.109 14.109 4 15"))),
		x.Child(x.Path(x.D("m6.5 12.5 1 1"))),
		x.Child(x.Path(x.D("m7 18 2.891 2.891"))),
		x.Child(x.Path(x.D("M9 22c1.798-1.998 2.518-3.995 2.807-5.993"))),
	)
	return x.Svg(svgArgs...)
}
