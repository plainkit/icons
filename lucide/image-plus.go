package lucide

import x "github.com/plainkit/html"

// ImagePlus creates a Image Plus Lucide icon.
func ImagePlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-image-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 5h6"))),
		x.Child(x.Path(x.D("M19 2v6"))),
		x.Child(x.Path(x.D("M21 11.5V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7.5"))),
		x.Child(x.Path(x.D("m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("9"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
