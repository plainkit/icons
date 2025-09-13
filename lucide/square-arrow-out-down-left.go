package lucide

import x "github.com/bloxui/blox"

// SquareArrowOutDownLeft creates a Square Arrow Out Down Left Lucide icon.
func SquareArrowOutDownLeft(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-out-down-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 21h6a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6"))),
		x.Child(x.Path(x.D("m3 21 9-9"))),
		x.Child(x.Path(x.D("M9 21H3v-6"))),
	)
	return x.Svg(svgArgs...)
}
