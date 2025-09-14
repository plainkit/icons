package lucide

import x "github.com/bloxui/blox"

// BanknoteX creates a Banknote X Lucide icon.
func BanknoteX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-banknote-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5"))),
		x.Child(x.Path(x.D("m17 17 5 5"))),
		x.Child(x.Path(x.D("M18 12h.01"))),
		x.Child(x.Path(x.D("m22 17-5 5"))),
		x.Child(x.Path(x.D("M6 12h.01"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
