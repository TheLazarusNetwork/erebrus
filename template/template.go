package template

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/TheLazarusNetwork/erebrus/model"
	"github.com/TheLazarusNetwork/erebrus/util"
)

var (
	wgTpl = `# Updated: {{ .Server.UpdatedAt }} / Created: {{ .Server.CreatedAt }}
[Interface]
{{- range .Server.Address }}
Address = {{ . }}
{{- end }}
ListenPort = {{ .Server.ListenPort }}
PrivateKey = {{ .Server.PrivateKey }}
{{ if ne .Server.Mtu 0 -}}
MTU = {{.Server.Mtu}}
{{- end}}
PreUp = {{ .Server.PreUp }}
PostUp = {{ .Server.PostUp }}
PreDown = {{ .Server.PreDown }}
PostDown = {{ .Server.PostDown }}
{{- range .Clients }}
{{ if .Enable -}}
# {{.Name}} / {{.Email}} / Updated: {{.UpdatedAt}} / Created: {{.CreatedAt}}
# friendly_name = {{.Name}}
[Peer]
PublicKey = {{ .PublicKey }}
PresharedKey = {{ .PresharedKey }}
AllowedIPs = {{ StringsJoin .Address ", " }}
{{- end }}
{{ end }}`
)

// DumpServerWg dump server wg config with go template, write it to file and return bytes
func DumpServerWg(clients []*model.Client, server *model.Server) ([]byte, error) {
	t, err := template.New("server").Funcs(template.FuncMap{"StringsJoin": strings.Join}).Parse(wgTpl)
	if err != nil {
		return nil, err
	}

	configDataWg, err := dump(t, struct {
		Clients []*model.Client
		Server  *model.Server
	}{
		Clients: clients,
		Server:  server,
	})
	if err != nil {
		return nil, err
	}

	err = util.WriteFile(filepath.Join(os.Getenv("WG_CONF_DIR"), os.Getenv("WG_INTERFACE_NAME")), configDataWg)
	if err != nil {
		return nil, err
	}

	return configDataWg, nil
}

func dump(tpl *template.Template, data interface{}) ([]byte, error) {
	var tplBuff bytes.Buffer

	err := tpl.Execute(&tplBuff, data)
	if err != nil {
		return nil, err
	}

	return tplBuff.Bytes(), nil
}

func FormatTime(t int64) string {
	result := time.Unix(0, t*int64(time.Millisecond))
	return result.Format("Monday, 02 January 06 15:04:05 MST")

}
