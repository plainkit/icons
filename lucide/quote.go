package lucide

import x "github.com/bloxui/blox"

// Quote creates a Quote Lucide icon.
func Quote(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-quote", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2 1 1 0 0 1 1 1v1a2 2 0 0 1-2 2 1 1 0 0 0-1 1v2a1 1 0 0 0 1 1 6 6 0 0 0 6-6V5a2 2 0 0 0-2-2z"))),
		x.Child(x.Path(x.D("M5 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2 1 1 0 0 1 1 1v1a2 2 0 0 1-2 2 1 1 0 0 0-1 1v2a1 1 0 0 0 1 1 6 6 0 0 0 6-6V5a2 2 0 0 0-2-2z"))),
	)
	return x.Svg(svgArgs...)
}
