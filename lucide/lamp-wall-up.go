package lucide

import x "github.com/bloxui/blox"

// LampWallUp creates a Lamp Wall Up Lucide icon.
func LampWallUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-lamp-wall-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19.929 9.629A1 1 0 0 1 19 11H9a1 1 0 0 1-.928-1.371l2-5A1 1 0 0 1 11 4h6a1 1 0 0 1 .928.629z"))),
		x.Child(x.Path(x.D("M6 15a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1z"))),
		x.Child(x.Path(x.D("M8 18h4a2 2 0 0 0 2-2v-5"))),
	)
	return x.Svg(svgArgs...)
}
