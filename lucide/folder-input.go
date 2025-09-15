package lucide

import x "github.com/plainkit/html"

// FolderInput creates a Folder Input Lucide icon.
func FolderInput(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-input", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 9V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-1"))),
		x.Child(x.Path(x.D("M2 13h10"))),
		x.Child(x.Path(x.D("m9 16 3-3-3-3"))),
	)
	return x.Svg(svgArgs...)
}
