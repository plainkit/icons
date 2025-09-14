package lucide

import x "github.com/bloxui/blox"

// FolderMinus creates a Folder Minus Lucide icon.
func FolderMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-folder-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 13h6"))),
		x.Child(x.Path(x.D("M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"))),
	)
	return x.Svg(svgArgs...)
}
