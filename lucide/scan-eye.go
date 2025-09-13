package lucide

import x "github.com/bloxui/blox"

// ScanEye creates a Scan Eye Lucide icon.
func ScanEye(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-scan-eye", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 7V5a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M17 3h2a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M21 17v2a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M7 21H5a2 2 0 0 1-2-2v-2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("1"))),
		x.Child(x.Path(x.D("M18.944 12.33a1 1 0 0 0 0-.66 7.5 7.5 0 0 0-13.888 0 1 1 0 0 0 0 .66 7.5 7.5 0 0 0 13.888 0"))),
	)
	return x.Svg(svgArgs...)
}
