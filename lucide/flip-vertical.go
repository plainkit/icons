package lucide

import x "github.com/bloxui/blox"

// FlipVertical creates a Flip Vertical Lucide icon.
func FlipVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flip-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 8V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v3"))),
		x.Child(x.Path(x.D("M21 16v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-3"))),
		x.Child(x.Path(x.D("M4 12H2"))),
		x.Child(x.Path(x.D("M10 12H8"))),
		x.Child(x.Path(x.D("M16 12h-2"))),
		x.Child(x.Path(x.D("M22 12h-2"))),
	)
	return x.Svg(svgArgs...)
}
