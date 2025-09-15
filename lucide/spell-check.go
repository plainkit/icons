package lucide

import x "github.com/plainkit/blox"

// SpellCheck creates a Spell Check Lucide icon.
func SpellCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-spell-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m6 16 6-12 6 12"))),
		x.Child(x.Path(x.D("M8 12h8"))),
		x.Child(x.Path(x.D("m16 20 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
