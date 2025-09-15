package lucide

import x "github.com/plainkit/blox"

// FolderSymlink creates a Folder Symlink Lucide icon.
func FolderSymlink(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-symlink", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 9.35V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h7"))),
		x.Child(x.Path(x.D("m8 16 3-3-3-3"))),
	)
	return x.Svg(svgArgs...)
}
