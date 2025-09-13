package lucide

import x "github.com/bloxui/blox"

// Heading creates a Heading Lucide icon.
func Heading(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-heading", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 12h12"))),
		x.Child(x.Path(x.D("M6 20V4"))),
		x.Child(x.Path(x.D("M18 20V4"))),
	)
	return x.Svg(svgArgs...)
}
