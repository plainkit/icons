package lucide

import x "github.com/bloxui/blox"

// FileSpreadsheet creates a File Spreadsheet Lucide icon.
func FileSpreadsheet(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-spreadsheet", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M8 13h2"))),
		x.Child(x.Path(x.D("M14 13h2"))),
		x.Child(x.Path(x.D("M8 17h2"))),
		x.Child(x.Path(x.D("M14 17h2"))),
	)
	return x.Svg(svgArgs...)
}
