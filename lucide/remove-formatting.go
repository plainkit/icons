package lucide

import x "github.com/plainkit/blox"

// RemoveFormatting creates a Remove Formatting Lucide icon.
func RemoveFormatting(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-remove-formatting", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 7V4h16v3"))),
		x.Child(x.Path(x.D("M5 20h6"))),
		x.Child(x.Path(x.D("M13 4 8 20"))),
		x.Child(x.Path(x.D("m15 15 5 5"))),
		x.Child(x.Path(x.D("m20 15-5 5"))),
	)
	return x.Svg(svgArgs...)
}
