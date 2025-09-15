package lucide

import x "github.com/plainkit/html"

// AlignCenterVertical creates a Align Center Vertical Lucide icon.
func AlignCenterVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-center-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v20"))),
		x.Child(x.Path(x.D("M8 10H4a2 2 0 0 1-2-2V6c0-1.1.9-2 2-2h4"))),
		x.Child(x.Path(x.D("M16 10h4a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-4"))),
		x.Child(x.Path(x.D("M8 20H7a2 2 0 0 1-2-2v-2c0-1.1.9-2 2-2h1"))),
		x.Child(x.Path(x.D("M16 14h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1"))),
	)
	return x.Svg(svgArgs...)
}
