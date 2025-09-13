package lucide

import x "github.com/bloxui/blox"

// HardDriveUpload creates a Hard Drive Upload Lucide icon.
func HardDriveUpload(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-hard-drive-upload", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 6-4-4-4 4"))),
		x.Child(x.Path(x.D("M12 2v8"))),
		x.Child(x.Rect(x.X("2"), x.Y("14"), x.RectWidth("20"), x.RectHeight("8"), x.Rx("2"))),
		x.Child(x.Path(x.D("M6 18h.01"))),
		x.Child(x.Path(x.D("M10 18h.01"))),
	)
	return x.Svg(svgArgs...)
}
