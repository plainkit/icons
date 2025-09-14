package lucide

import x "github.com/bloxui/blox"

// Space creates a Space Lucide icon.
func Space(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-space", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 17v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1"))),
	)
	return x.Svg(svgArgs...)
}
