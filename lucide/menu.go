package lucide

import x "github.com/bloxui/blox"

// Menu creates a Menu Lucide icon.
func Menu(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-menu", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 5h16"))),
		x.Child(x.Path(x.D("M4 12h16"))),
		x.Child(x.Path(x.D("M4 19h16"))),
	)
	return x.Svg(svgArgs...)
}
