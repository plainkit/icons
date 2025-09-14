package lucide

import x "github.com/bloxui/blox"

// Import creates a Import Lucide icon.
func Import(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-import", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3v12"))),
		x.Child(x.Path(x.D("m8 11 4 4 4-4"))),
		x.Child(x.Path(x.D("M8 5H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-4"))),
	)
	return x.Svg(svgArgs...)
}
