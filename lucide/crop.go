package lucide

import x "github.com/bloxui/blox"

// Crop creates a Crop Lucide icon.
func Crop(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-crop", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 2v14a2 2 0 0 0 2 2h14"))),
		x.Child(x.Path(x.D("M18 22V8a2 2 0 0 0-2-2H2"))),
	)
	return x.Svg(svgArgs...)
}
