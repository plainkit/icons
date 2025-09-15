package lucide

import x "github.com/plainkit/html"

// Brain creates a Brain Lucide icon.
func Brain(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-brain", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 18V5"))),
		x.Child(x.Path(x.D("M15 13a4.17 4.17 0 0 1-3-4 4.17 4.17 0 0 1-3 4"))),
		x.Child(x.Path(x.D("M17.598 6.5A3 3 0 1 0 12 5a3 3 0 1 0-5.598 1.5"))),
		x.Child(x.Path(x.D("M17.997 5.125a4 4 0 0 1 2.526 5.77"))),
		x.Child(x.Path(x.D("M18 18a4 4 0 0 0 2-7.464"))),
		x.Child(x.Path(x.D("M19.967 17.483A4 4 0 1 1 12 18a4 4 0 1 1-7.967-.517"))),
		x.Child(x.Path(x.D("M6 18a4 4 0 0 1-2-7.464"))),
		x.Child(x.Path(x.D("M6.003 5.125a4 4 0 0 0-2.526 5.77"))),
	)
	return x.Svg(svgArgs...)
}
