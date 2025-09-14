package lucide

import x "github.com/bloxui/blox"

// SquareStack creates a Square Stack Lucide icon.
func SquareStack(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-stack", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 10c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2"))),
		x.Child(x.Path(x.D("M10 16c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("8"), x.X("14"), x.Y("14"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
