package lucide

import x "github.com/bloxui/blox"

// FolderRoot creates a Folder Root Lucide icon.
func FolderRoot(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-folder-root", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("13"), x.R("2"))),
		x.Child(x.Path(x.D("M12 15v5"))),
	)
	return x.Svg(svgArgs...)
}
