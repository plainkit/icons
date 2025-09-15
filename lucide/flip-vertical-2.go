package lucide

import x "github.com/plainkit/blox"

// FlipVertical2 creates a Flip Vertical 2 Lucide icon.
func FlipVertical2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flip-vertical-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m17 3-5 5-5-5h10"))),
		x.Child(x.Path(x.D("m17 21-5-5-5 5h10"))),
		x.Child(x.Path(x.D("M4 12H2"))),
		x.Child(x.Path(x.D("M10 12H8"))),
		x.Child(x.Path(x.D("M16 12h-2"))),
		x.Child(x.Path(x.D("M22 12h-2"))),
	)
	return x.Svg(svgArgs...)
}
