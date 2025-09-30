package lucide

import (
	html "github.com/plainkit/html"
)

// FileQuestionMark creates a File Question Mark Lucide icon.
func FileQuestionMark(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-question-mark", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 17h.01")),
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7z")),
		html.SvgPath(html.AD("M9.1 9a3 3 0 0 1 5.82 1c0 2-3 3-3 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
