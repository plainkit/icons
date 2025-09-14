package lucide

import x "github.com/bloxui/blox"

// Triangle creates a Triangle Lucide icon.
func Triangle(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-triangle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.73 4a2 2 0 0 0-3.46 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3Z"))),
	)
	return x.Svg(svgArgs...)
}
