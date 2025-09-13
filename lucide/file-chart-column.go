package lucide

import x "github.com/bloxui/blox"

// FileChartColumn creates a File Chart Column Lucide icon.
func FileChartColumn(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-chart-column", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M8 18v-1"))),
		x.Child(x.Path(x.D("M12 18v-6"))),
		x.Child(x.Path(x.D("M16 18v-3"))),
	)
	return x.Svg(svgArgs...)
}
