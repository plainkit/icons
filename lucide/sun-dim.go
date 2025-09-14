package lucide

import x "github.com/bloxui/blox"

// SunDim creates a Sun Dim Lucide icon.
func SunDim(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sun-dim", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("4"))),
		x.Child(x.Path(x.D("M12 4h.01"))),
		x.Child(x.Path(x.D("M20 12h.01"))),
		x.Child(x.Path(x.D("M12 20h.01"))),
		x.Child(x.Path(x.D("M4 12h.01"))),
		x.Child(x.Path(x.D("M17.657 6.343h.01"))),
		x.Child(x.Path(x.D("M17.657 17.657h.01"))),
		x.Child(x.Path(x.D("M6.343 17.657h.01"))),
		x.Child(x.Path(x.D("M6.343 6.343h.01"))),
	)
	return x.Svg(svgArgs...)
}
