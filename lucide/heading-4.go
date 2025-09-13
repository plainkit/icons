package lucide

import x "github.com/bloxui/blox"

// Heading4 creates a Heading 4 Lucide icon.
func Heading4(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-heading-4", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 18V6"))),
		x.Child(x.Path(x.D("M17 10v3a1 1 0 0 0 1 1h3"))),
		x.Child(x.Path(x.D("M21 10v8"))),
		x.Child(x.Path(x.D("M4 12h8"))),
		x.Child(x.Path(x.D("M4 18V6"))),
	)
	return x.Svg(svgArgs...)
}
