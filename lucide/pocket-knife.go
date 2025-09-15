package lucide

import x "github.com/plainkit/blox"

// PocketKnife creates a Pocket Knife Lucide icon.
func PocketKnife(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pocket-knife", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 2v1c0 1 2 1 2 2S3 6 3 7s2 1 2 2-2 1-2 2 2 1 2 2"))),
		x.Child(x.Path(x.D("M18 6h.01"))),
		x.Child(x.Path(x.D("M6 18h.01"))),
		x.Child(x.Path(x.D("M20.83 8.83a4 4 0 0 0-5.66-5.66l-12 12a4 4 0 1 0 5.66 5.66Z"))),
		x.Child(x.Path(x.D("M18 11.66V22a4 4 0 0 0 4-4V6"))),
	)
	return x.Svg(svgArgs...)
}
