package lucide

import x "github.com/plainkit/blox"

// Headset creates a Headset Lucide icon.
func Headset(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-headset", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 11h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-5Zm0 0a9 9 0 1 1 18 0m0 0v5a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3Z"))),
		x.Child(x.Path(x.D("M21 16v2a4 4 0 0 1-4 4h-5"))),
	)
	return x.Svg(svgArgs...)
}
