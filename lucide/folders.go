package lucide

import x "github.com/bloxui/blox"

// Folders creates a Folders Lucide icon.
func Folders(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-folders", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 5a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h2.5a1.5 1.5 0 0 1 1.2.6l.6.8a1.5 1.5 0 0 0 1.2.6z"))),
		x.Child(x.Path(x.D("M3 8.268a2 2 0 0 0-1 1.738V19a2 2 0 0 0 2 2h11a2 2 0 0 0 1.732-1"))),
	)
	return x.Svg(svgArgs...)
}
