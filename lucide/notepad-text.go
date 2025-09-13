package lucide

import x "github.com/bloxui/blox"

// NotepadText creates a Notepad Text Lucide icon.
func NotepadText(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-notepad-text", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M12 2v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("18"), x.X("4"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M8 10h6"))),
		x.Child(x.Path(x.D("M8 14h8"))),
		x.Child(x.Path(x.D("M8 18h5"))),
	)
	return x.Svg(svgArgs...)
}
