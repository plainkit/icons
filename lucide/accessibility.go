package lucide

import x "github.com/bloxui/blox"

// Accessibility creates a Accessibility Lucide icon.
func Accessibility(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-accessibility", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("16"), x.Cy("4"), x.R("1"))),
		x.Child(x.Path(x.D("m18 19 1-7-6 1"))),
		x.Child(x.Path(x.D("m5 8 3-3 5.5 3-2.36 3.5"))),
		x.Child(x.Path(x.D("M4.24 14.5a5 5 0 0 0 6.88 6"))),
		x.Child(x.Path(x.D("M13.76 17.5a5 5 0 0 0-6.88-6"))),
	)
	return x.Svg(svgArgs...)
}
