package lucide

import x "github.com/bloxui/blox"

// FileTerminal creates a File Terminal Lucide icon.
func FileTerminal(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-terminal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("m8 16 2-2-2-2"))),
		x.Child(x.Path(x.D("M12 18h4"))),
	)
	return x.Svg(svgArgs...)
}
