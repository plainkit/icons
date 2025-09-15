package lucide

import x "github.com/plainkit/html"

// BanknoteArrowUp creates a Banknote Arrow Up Lucide icon.
func BanknoteArrowUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-banknote-arrow-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5"))),
		x.Child(x.Path(x.D("M18 12h.01"))),
		x.Child(x.Path(x.D("M19 22v-6"))),
		x.Child(x.Path(x.D("m22 19-3-3-3 3"))),
		x.Child(x.Path(x.D("M6 12h.01"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
