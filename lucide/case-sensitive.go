package lucide

import x "github.com/bloxui/blox"

// CaseSensitive creates a Case Sensitive Lucide icon.
func CaseSensitive(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-case-sensitive", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16"))),
		x.Child(x.Path(x.D("M22 9v7"))),
		x.Child(x.Path(x.D("M3.304 13h6.392"))),
		x.Child(x.Circle(x.Cx("18.5"), x.Cy("12.5"), x.R("3.5"))),
	)
	return x.Svg(svgArgs...)
}
