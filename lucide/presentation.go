package lucide

import x "github.com/bloxui/blox"

// Presentation creates a Presentation Lucide icon.
func Presentation(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-presentation", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 3h20"))),
		x.Child(x.Path(x.D("M21 3v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V3"))),
		x.Child(x.Path(x.D("m7 21 5-5 5 5"))),
	)
	return x.Svg(svgArgs...)
}
