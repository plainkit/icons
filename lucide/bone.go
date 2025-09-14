package lucide

import x "github.com/bloxui/blox"

// Bone creates a Bone Lucide icon.
func Bone(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bone", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 10c.7-.7 1.69 0 2.5 0a2.5 2.5 0 1 0 0-5 .5.5 0 0 1-.5-.5 2.5 2.5 0 1 0-5 0c0 .81.7 1.8 0 2.5l-7 7c-.7.7-1.69 0-2.5 0a2.5 2.5 0 0 0 0 5c.28 0 .5.22.5.5a2.5 2.5 0 1 0 5 0c0-.81-.7-1.8 0-2.5Z"))),
	)
	return x.Svg(svgArgs...)
}
