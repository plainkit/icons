package lucide

import x "github.com/plainkit/blox"

// Download creates a Download Lucide icon.
func Download(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-download", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 15V3"))),
		x.Child(x.Path(x.D("M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"))),
		x.Child(x.Path(x.D("m7 10 5 5 5-5"))),
	)
	return x.Svg(svgArgs...)
}
