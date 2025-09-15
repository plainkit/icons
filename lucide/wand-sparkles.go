package lucide

import x "github.com/plainkit/html"

// WandSparkles creates a Wand Sparkles Lucide icon.
func WandSparkles(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wand-sparkles", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m21.64 3.64-1.28-1.28a1.21 1.21 0 0 0-1.72 0L2.36 18.64a1.21 1.21 0 0 0 0 1.72l1.28 1.28a1.2 1.2 0 0 0 1.72 0L21.64 5.36a1.2 1.2 0 0 0 0-1.72"))),
		x.Child(x.Path(x.D("m14 7 3 3"))),
		x.Child(x.Path(x.D("M5 6v4"))),
		x.Child(x.Path(x.D("M19 14v4"))),
		x.Child(x.Path(x.D("M10 2v2"))),
		x.Child(x.Path(x.D("M7 8H3"))),
		x.Child(x.Path(x.D("M21 16h-4"))),
		x.Child(x.Path(x.D("M11 3H9"))),
	)
	return x.Svg(svgArgs...)
}
