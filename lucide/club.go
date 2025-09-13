package lucide

import x "github.com/bloxui/blox"

// Club creates a Club Lucide icon.
func Club(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-club", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17.28 9.05a5.5 5.5 0 1 0-10.56 0A5.5 5.5 0 1 0 12 17.66a5.5 5.5 0 1 0 5.28-8.6Z"))),
		x.Child(x.Path(x.D("M12 17.66L12 22"))),
	)
	return x.Svg(svgArgs...)
}
