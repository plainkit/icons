package lucide

import x "github.com/bloxui/blox"

// Asterisk creates an Asterisk Lucide icon.
func Asterisk(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-asterisk", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v12"))),
		x.Child(x.Path(x.D("M17.196 9 6.804 15"))),
		x.Child(x.Path(x.D("m6.804 9 10.392 6"))),
	)
	return x.Svg(svgArgs...)
}
