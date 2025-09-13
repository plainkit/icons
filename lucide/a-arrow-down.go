package lucide

import x "github.com/bloxui/blox"

// AArrowDown creates an A Arrow Down Lucide icon
func AArrowDown(args ...x.SvgArg) x.Component {
	// Build SVG arguments with proper override handling
	svgArgs := buildLucideArgs("lucide lucide-a-arrow-down", args)

	// Add the icon paths
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14 12 4 4 4-4"))),
		x.Child(x.Path(x.D("M18 16V7"))),
		x.Child(x.Path(x.D("m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16"))),
		x.Child(x.Path(x.D("M3.304 13h6.392"))),
	)

	return x.Svg(svgArgs...)
}
