package lucide

import x "github.com/bloxui/blox"

// FolderArchive creates a Folder Archive Lucide icon.
func FolderArchive(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-archive", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("15"), x.Cy("19"), x.R("2"))),
		x.Child(x.Path(x.D("M20.9 19.8A2 2 0 0 0 22 18V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h5.1"))),
		x.Child(x.Path(x.D("M15 11v-1"))),
		x.Child(x.Path(x.D("M15 17v-2"))),
	)
	return x.Svg(svgArgs...)
}
