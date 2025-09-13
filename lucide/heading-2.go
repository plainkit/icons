package lucide

import x "github.com/bloxui/blox"

// Heading2 creates a Heading 2 Lucide icon.
func Heading2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-heading-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 12h8"))),
		x.Child(x.Path(x.D("M4 18V6"))),
		x.Child(x.Path(x.D("M12 18V6"))),
		x.Child(x.Path(x.D("M21 18h-4c0-4 4-3 4-6 0-1.5-2-2.5-4-1"))),
	)
	return x.Svg(svgArgs...)
}
