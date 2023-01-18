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
	clientTpl = `[Interface]
Address = {{ StringsJoin .Client.Address ", " }}
PrivateKey = {{ .Client.PrivateKey }}
{{ if ne (len .Server.DNS) 0 -}}
DNS = {{ StringsJoin .Server.DNS ", " }}
{{- end }}
{{ if ne .Server.Mtu 0 -}}
MTU = {{.Server.Mtu}}
{{- end}}
[Peer]
PublicKey = {{ .Server.PublicKey }}
PresharedKey = {{ .Client.PresharedKey }}
AllowedIPs = {{ StringsJoin .Client.AllowedIPs ", " }}
Endpoint = {{ .Server.Endpoint }}:{{ .Server.ListenPort }}
{{ if and (ne .Server.PersistentKeepalive 0) (not .Client.IgnorePersistentKeepalive) -}}
PersistentKeepalive = {{.Server.PersistentKeepalive}}
{{- end}}
`

	wgTpl = `# Updated: {{ .Server.Updated }} / Created: {{ .Server.Created }}
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
# {{.Name}} / {{.Email}} / Updated: {{.Updated}} / Created: {{.Created}}
# friendly_name = {{.Name}}
[Peer]
PublicKey = {{ .PublicKey }}
PresharedKey = {{ .PresharedKey }}
AllowedIPs = {{ StringsJoin .Address ", " }}
{{- end }}
{{ end }}`
)

// DumpClientWg dump client wg config with go template
func DumpClientWg(client *model.Client, server *model.Server) ([]byte, error) {
	t, err := template.New("client").Funcs(template.FuncMap{"StringsJoin": strings.Join}).Parse(clientTpl)
	if err != nil {
		return nil, err
	}

	return dump(t, struct {
		Client *model.Client
		Server *model.Server
	}{
		Client: client,
		Server: server,
	})
}

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
