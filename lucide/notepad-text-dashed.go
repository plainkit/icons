package lucide

import x "github.com/plainkit/html"

// NotepadTextDashed creates a Notepad Text Dashed Lucide icon.
func NotepadTextDashed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-notepad-text-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M12 2v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M16 4h2a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M20 12v2"))),
		x.Child(x.Path(x.D("M20 18v2a2 2 0 0 1-2 2h-1"))),
		x.Child(x.Path(x.D("M13 22h-2"))),
		x.Child(x.Path(x.D("M7 22H6a2 2 0 0 1-2-2v-2"))),
		x.Child(x.Path(x.D("M4 14v-2"))),
		x.Child(x.Path(x.D("M4 8V6a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M8 10h6"))),
		x.Child(x.Path(x.D("M8 14h8"))),
		x.Child(x.Path(x.D("M8 18h5"))),
	)
	return x.Svg(svgArgs...)
}
