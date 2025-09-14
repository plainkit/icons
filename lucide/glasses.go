package lucide

import x "github.com/bloxui/blox"

// Glasses creates a Glasses Lucide icon.
func Glasses(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-glasses", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("6"), x.Cy("15"), x.R("4"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("15"), x.R("4"))),
		x.Child(x.Path(x.D("M14 15a2 2 0 0 0-2-2 2 2 0 0 0-2 2"))),
		x.Child(x.Path(x.D("M2.5 13 5 7c.7-1.3 1.4-2 3-2"))),
		x.Child(x.Path(x.D("M21.5 13 19 7c-.7-1.3-1.5-2-3-2"))),
	)
	return x.Svg(svgArgs...)
}
