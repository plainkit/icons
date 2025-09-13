package lucide

import x "github.com/bloxui/blox"

// SkipBack creates a Skip Back Lucide icon.
func SkipBack(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-skip-back", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17.971 4.285A2 2 0 0 1 21 6v12a2 2 0 0 1-3.029 1.715l-9.997-5.998a2 2 0 0 1-.003-3.432z"))),
		x.Child(x.Path(x.D("M3 20V4"))),
	)
	return x.Svg(svgArgs...)
}
