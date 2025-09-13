package lucide

import x "github.com/bloxui/blox"

// FileChartLine creates a File Chart Line Lucide icon.
func FileChartLine(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-chart-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("m16 13-3.5 3.5-2-2L8 17"))),
	)
	return x.Svg(svgArgs...)
}