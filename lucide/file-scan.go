package lucide

import x "github.com/bloxui/blox"

// FileScan creates a File Scan Lucide icon.
func FileScan(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-scan", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 10V7l-5-5H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M16 14a2 2 0 0 0-2 2"))),
		x.Child(x.Path(x.D("M20 14a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M20 22a2 2 0 0 0 2-2"))),
		x.Child(x.Path(x.D("M16 22a2 2 0 0 1-2-2"))),
	)
	return x.Svg(svgArgs...)
}
