package lucide

import x "github.com/bloxui/blox"

// DnaOff creates a Dna Off Lucide icon.
func DnaOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-dna-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2c-1.35 1.5-2.092 3-2.5 4.5L14 8"))),
		x.Child(x.Path(x.D("m17 6-2.891-2.891"))),
		x.Child(x.Path(x.D("M2 15c3.333-3 6.667-3 10-3"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("m20 9 .891.891"))),
		x.Child(x.Path(x.D("M22 9c-1.5 1.35-3 2.092-4.5 2.5l-1-1"))),
		x.Child(x.Path(x.D("M3.109 14.109 4 15"))),
		x.Child(x.Path(x.D("m6.5 12.5 1 1"))),
		x.Child(x.Path(x.D("m7 18 2.891 2.891"))),
		x.Child(x.Path(x.D("M9 22c1.35-1.5 2.092-3 2.5-4.5L10 16"))),
	)
	return x.Svg(svgArgs...)
}
