package lucide

import x "github.com/bloxui/blox"

// SeparatorVertical creates a Separator Vertical Lucide icon.
func SeparatorVertical(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-separator-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3v18"))),
		x.Child(x.Path(x.D("m16 16 4-4-4-4"))),
		x.Child(x.Path(x.D("m8 8-4 4 4 4"))),
	)
	return x.Svg(svgArgs...)
}
