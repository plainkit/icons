package lucide

import x "github.com/bloxui/blox"

// Save creates a Save Lucide icon.
func Save(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-save", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15.2 3a2 2 0 0 1 1.4.6l3.8 3.8a2 2 0 0 1 .6 1.4V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2z"))),
		x.Child(x.Path(x.D("M17 21v-7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v7"))),
		x.Child(x.Path(x.D("M7 3v4a1 1 0 0 0 1 1h7"))),
	)
	return x.Svg(svgArgs...)
}
