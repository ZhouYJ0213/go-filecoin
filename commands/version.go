package commands

import (
	"fmt"
	"io"

	cmds "gx/ipfs/QmQtQrtNioesAWtrx8csBvfY37gTe94d6wQ3VikZUjxD39/go-ipfs-cmds"
	cmdkit "gx/ipfs/Qmde5VP1qUkyQXKCfmEUA7bP64V2HAptbJ7phuPp7jXWwg/go-ipfs-cmdkit"

	"github.com/filecoin-project/go-filecoin/flags"
)

type versionInfo struct {
	// Commit, is the git sha that was used to build this version of go-filecoin.
	Commit string
}

var versionCmd = &cmds.Command{
	Helptext: cmdkit.HelpText{
		Tagline: "Show go-filecoin version information",
	},
	Run: func(req *cmds.Request, re cmds.ResponseEmitter, env cmds.Environment) error {
		return re.Emit(&versionInfo{
			Commit: flags.Commit,
		})
	},
	Type: versionInfo{},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request, w io.Writer, vo *versionInfo) error {
			_, err := fmt.Fprintf(w, "commit: %s\n", vo.Commit)
			return err
		}),
	},
}
