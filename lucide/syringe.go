package lucide

import x "github.com/bloxui/blox"

// Syringe creates a Syringe Lucide icon.
func Syringe(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-syringe", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m18 2 4 4"))),
		x.Child(x.Path(x.D("m17 7 3-3"))),
		x.Child(x.Path(x.D("M19 9 8.7 19.3c-1 1-2.5 1-3.4 0l-.6-.6c-1-1-1-2.5 0-3.4L15 5"))),
		x.Child(x.Path(x.D("m9 11 4 4"))),
		x.Child(x.Path(x.D("m5 19-3 3"))),
		x.Child(x.Path(x.D("m14 4 6 6"))),
	)
	return x.Svg(svgArgs...)
}
