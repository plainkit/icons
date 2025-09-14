package lucide

import x "github.com/bloxui/blox"

// FlipHorizontal creates a Flip Horizontal Lucide icon.
func FlipHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flip-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h3"))),
		x.Child(x.Path(x.D("M16 3h3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-3"))),
		x.Child(x.Path(x.D("M12 20v2"))),
		x.Child(x.Path(x.D("M12 14v2"))),
		x.Child(x.Path(x.D("M12 8v2"))),
		x.Child(x.Path(x.D("M12 2v2"))),
	)
	return x.Svg(svgArgs...)
}
