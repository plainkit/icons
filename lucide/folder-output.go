package lucide

import x "github.com/bloxui/blox"

// FolderOutput creates a Folder Output Lucide icon.
func FolderOutput(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-output", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 7.5V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-1.5"))),
		x.Child(x.Path(x.D("M2 13h10"))),
		x.Child(x.Path(x.D("m5 10-3 3 3 3"))),
	)
	return x.Svg(svgArgs...)
}
