package lucide

import x "github.com/bloxui/blox"

// Share creates a Share Lucide icon.
func Share(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-share", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v13"))),
		x.Child(x.Path(x.D("m16 6-4-4-4 4"))),
		x.Child(x.Path(x.D("M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"))),
	)
	return x.Svg(svgArgs...)
}
