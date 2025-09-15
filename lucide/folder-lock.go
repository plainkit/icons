package lucide

import x "github.com/plainkit/html"

// FolderLock creates a Folder Lock Lucide icon.
func FolderLock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-lock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("5"), x.X("14"), x.Y("17"), x.Rx("1"))),
		x.Child(x.Path(x.D("M10 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v2.5"))),
		x.Child(x.Path(x.D("M20 17v-2a2 2 0 1 0-4 0v2"))),
	)
	return x.Svg(svgArgs...)
}
