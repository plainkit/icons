package lucide

import x "github.com/bloxui/blox"

// Pipette creates a Pipette Lucide icon.
func Pipette(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-pipette", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12 9-8.414 8.414A2 2 0 0 0 3 18.828v1.344a2 2 0 0 1-.586 1.414A2 2 0 0 1 3.828 21h1.344a2 2 0 0 0 1.414-.586L15 12"))),
		x.Child(x.Path(x.D("m18 9 .4.4a1 1 0 1 1-3 3l-3.8-3.8a1 1 0 1 1 3-3l.4.4 3.4-3.4a1 1 0 1 1 3 3z"))),
		x.Child(x.Path(x.D("m2 22 .414-.414"))),
	)
	return x.Svg(svgArgs...)
}
