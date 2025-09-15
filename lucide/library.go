package lucide

import x "github.com/plainkit/blox"

// Library creates a Library Lucide icon.
func Library(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-library", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 6 4 14"))),
		x.Child(x.Path(x.D("M12 6v14"))),
		x.Child(x.Path(x.D("M8 8v12"))),
		x.Child(x.Path(x.D("M4 4v16"))),
	)
	return x.Svg(svgArgs...)
}
