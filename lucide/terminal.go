package lucide

import x "github.com/bloxui/blox"

// Terminal creates a Terminal Lucide icon.
func Terminal(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-terminal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 19h8"))),
		x.Child(x.Path(x.D("m4 17 6-6-6-6"))),
	)
	return x.Svg(svgArgs...)
}
