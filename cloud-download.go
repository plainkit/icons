package lucide

import x "github.com/bloxui/blox"

// CloudDownload creates a Cloud Download Lucide icon.
func CloudDownload(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-cloud-download", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 13v8l-4-4"))),
		x.Child(x.Path(x.D("m12 21 4-4"))),
		x.Child(x.Path(x.D("M4.393 15.269A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.436 8.284"))),
	)
	return x.Svg(svgArgs...)
}
