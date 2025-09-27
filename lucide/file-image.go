package lucide

import (
	html "github.com/plainkit/html"
)

// FileImage creates a File Image Lucide icon.
func FileImage(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-image", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("12"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("m20 17-1.296-1.296a2.41 2.41 0 0 0-3.408 0L9 22"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
