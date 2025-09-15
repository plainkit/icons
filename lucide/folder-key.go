package lucide

import x "github.com/plainkit/blox"

// FolderKey creates a Folder Key Lucide icon.
func FolderKey(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-key", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("16"), x.Cy("20"), x.R("2"))),
		x.Child(x.Path(x.D("M10 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("m22 14-4.5 4.5"))),
		x.Child(x.Path(x.D("m21 15 1 1"))),
	)
	return x.Svg(svgArgs...)
}
