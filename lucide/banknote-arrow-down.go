package lucide

import x "github.com/bloxui/blox"

// BanknoteArrowDown creates a Banknote Arrow Down Lucide icon.
func BanknoteArrowDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-banknote-arrow-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5"))),
		x.Child(x.Path(x.D("m16 19 3 3 3-3"))),
		x.Child(x.Path(x.D("M18 12h.01"))),
		x.Child(x.Path(x.D("M19 16v6"))),
		x.Child(x.Path(x.D("M6 12h.01"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
