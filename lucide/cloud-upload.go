package lucide

import x "github.com/bloxui/blox"

// CloudUpload creates a Cloud Upload Lucide icon.
func CloudUpload(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-cloud-upload", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 13v8"))),
		x.Child(x.Path(x.D("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"))),
		x.Child(x.Path(x.D("m8 17 4-4 4 4"))),
	)
	return x.Svg(svgArgs...)
}
