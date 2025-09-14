package lucide

import x "github.com/bloxui/blox"

// FlipHorizontal2 creates a Flip Horizontal 2 Lucide icon.
func FlipHorizontal2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flip-horizontal-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 7 5 5-5 5V7"))),
		x.Child(x.Path(x.D("m21 7-5 5 5 5V7"))),
		x.Child(x.Path(x.D("M12 20v2"))),
		x.Child(x.Path(x.D("M12 14v2"))),
		x.Child(x.Path(x.D("M12 8v2"))),
		x.Child(x.Path(x.D("M12 2v2"))),
	)
	return x.Svg(svgArgs...)
}
