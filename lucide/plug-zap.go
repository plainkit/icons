package lucide

import x "github.com/bloxui/blox"

// PlugZap creates a Plug Zap Lucide icon.
func PlugZap(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-plug-zap", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6.3 20.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6-2.3 2.3a2.4 2.4 0 0 0 0 3.4Z"))),
		x.Child(x.Path(x.D("m2 22 3-3"))),
		x.Child(x.Path(x.D("M7.5 13.5 10 11"))),
		x.Child(x.Path(x.D("M10.5 16.5 13 14"))),
		x.Child(x.Path(x.D("m18 3-4 4h6l-4 4"))),
	)
	return x.Svg(svgArgs...)
}
