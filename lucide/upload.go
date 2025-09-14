package lucide

import x "github.com/bloxui/blox"

// Upload creates a Upload Lucide icon.
func Upload(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-upload", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3v12"))),
		x.Child(x.Path(x.D("m17 8-5-5-5 5"))),
		x.Child(x.Path(x.D("M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"))),
	)
	return x.Svg(svgArgs...)
}
