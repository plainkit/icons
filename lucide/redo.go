package lucide

import x "github.com/bloxui/blox"

// Redo creates a Redo Lucide icon.
func Redo(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-redo", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 7v6h-6"))),
		x.Child(x.Path(x.D("M3 17a9 9 0 0 1 9-9 9 9 0 0 1 6 2.3l3 2.7"))),
	)
	return x.Svg(svgArgs...)
}
