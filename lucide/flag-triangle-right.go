package lucide

import x "github.com/bloxui/blox"

// FlagTriangleRight creates a Flag Triangle Right Lucide icon.
func FlagTriangleRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flag-triangle-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 22V2.8a.8.8 0 0 1 1.17-.71l11.38 5.69a.8.8 0 0 1 0 1.44L6 15.5"))),
	)
	return x.Svg(svgArgs...)
}
