package lucide

import x "github.com/bloxui/blox"

// ArchiveRestore creates a Archive Restore Lucide icon.
func ArchiveRestore(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-archive-restore", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("5"), x.X("2"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Path(x.D("M4 8v11a2 2 0 0 0 2 2h2"))),
		x.Child(x.Path(x.D("M20 8v11a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("m9 15 3-3 3 3"))),
		x.Child(x.Path(x.D("M12 12v9"))),
	)
	return x.Svg(svgArgs...)
}
