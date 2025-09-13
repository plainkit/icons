package lucide

import x "github.com/bloxui/blox"

// SpellCheck2 creates a Spell Check 2 Lucide icon.
func SpellCheck2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-spell-check-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m6 16 6-12 6 12"))),
		x.Child(x.Path(x.D("M8 12h8"))),
		x.Child(x.Path(x.D("M4 21c1.1 0 1.1-1 2.3-1s1.1 1 2.3 1c1.1 0 1.1-1 2.3-1 1.1 0 1.1 1 2.3 1 1.1 0 1.1-1 2.3-1 1.1 0 1.1 1 2.3 1 1.1 0 1.1-1 2.3-1"))),
	)
	return x.Svg(svgArgs...)
}
