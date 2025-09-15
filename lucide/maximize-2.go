package lucide

import x "github.com/plainkit/html"

// Maximize2 creates a Maximize 2 Lucide icon.
func Maximize2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-maximize-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 3h6v6"))),
		x.Child(x.Path(x.D("m21 3-7 7"))),
		x.Child(x.Path(x.D("m3 21 7-7"))),
		x.Child(x.Path(x.D("M9 21H3v-6"))),
	)
	return x.Svg(svgArgs...)
}
