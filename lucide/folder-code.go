package lucide

import x "github.com/bloxui/blox"

// FolderCode creates a Folder Code Lucide icon.
func FolderCode(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-folder-code", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 10.5 8 13l2 2.5"))),
		x.Child(x.Path(x.D("m14 10.5 2 2.5-2 2.5"))),
		x.Child(x.Path(x.D("M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2z"))),
	)
	return x.Svg(svgArgs...)
}
