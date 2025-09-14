package lucide

import x "github.com/bloxui/blox"

// FolderSync creates a Folder Sync Lucide icon.
func FolderSync(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-sync", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v.5"))),
		x.Child(x.Path(x.D("M12 10v4h4"))),
		x.Child(x.Path(x.D("m12 14 1.535-1.605a5 5 0 0 1 8 1.5"))),
		x.Child(x.Path(x.D("M22 22v-4h-4"))),
		x.Child(x.Path(x.D("m22 18-1.535 1.605a5 5 0 0 1-8-1.5"))),
	)
	return x.Svg(svgArgs...)
}
