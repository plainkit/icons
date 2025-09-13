package lucide

import x "github.com/bloxui/blox"

// Command creates a Command Lucide icon.
func Command(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-command", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 6v12a3 3 0 1 0 3-3H6a3 3 0 1 0 3 3V6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3"))),
	)
	return x.Svg(svgArgs...)
}
