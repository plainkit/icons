package lucide

import x "github.com/plainkit/blox"

// NonBinary creates a Non Binary Lucide icon.
func NonBinary(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-non-binary", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v10"))),
		x.Child(x.Path(x.D("m8.5 4 7 4"))),
		x.Child(x.Path(x.D("m8.5 8 7-4"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("17"), x.R("5"))),
	)
	return x.Svg(svgArgs...)
}
