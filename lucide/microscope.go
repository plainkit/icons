package lucide

import x "github.com/bloxui/blox"

// Microscope creates a Microscope Lucide icon.
func Microscope(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-microscope", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 18h8"))),
		x.Child(x.Path(x.D("M3 22h18"))),
		x.Child(x.Path(x.D("M14 22a7 7 0 1 0 0-14h-1"))),
		x.Child(x.Path(x.D("M9 14h2"))),
		x.Child(x.Path(x.D("M9 12a2 2 0 0 1-2-2V6h6v4a2 2 0 0 1-2 2Z"))),
		x.Child(x.Path(x.D("M12 6V3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3"))),
	)
	return x.Svg(svgArgs...)
}
