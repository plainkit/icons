package lucide

import x "github.com/bloxui/blox"

// Stamp creates a Stamp Lucide icon.
func Stamp(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-stamp", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 13V8.5C14 7 15 7 15 5a3 3 0 0 0-6 0c0 2 1 2 1 3.5V13"))),
		x.Child(x.Path(x.D("M20 15.5a2.5 2.5 0 0 0-2.5-2.5h-11A2.5 2.5 0 0 0 4 15.5V17a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1z"))),
		x.Child(x.Path(x.D("M5 22h14"))),
	)
	return x.Svg(svgArgs...)
}
