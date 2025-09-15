package lucide

import x "github.com/plainkit/html"

// Shredder creates a Shredder Lucide icon.
func Shredder(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shredder", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 22v-5"))),
		x.Child(x.Path(x.D("M14 19v-2"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M18 20v-3"))),
		x.Child(x.Path(x.D("M2 13h20"))),
		x.Child(x.Path(x.D("M20 13V7l-5-5H6a2 2 0 0 0-2 2v9"))),
		x.Child(x.Path(x.D("M6 20v-3"))),
	)
	return x.Svg(svgArgs...)
}
