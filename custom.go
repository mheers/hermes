package hermes

// Custom is the theme by default
type Custom struct{}

// Name returns the name of the default theme
func (dt *Custom) Name() string {
	return "custom"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *Custom) HTMLTemplate() string {
	return `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <style type="text/css" rel="stylesheet" media="all">
    /* Base ------------------------------ */
    *:not(br):not(tr):not(html) {
      font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;
      -webkit-box-sizing: border-box;
      box-sizing: border-box;
    }
    body {
      width: 100% !important;
      height: 100%;
      margin: 0;
      line-height: 1.4;
      background-color: {{ .Hermes.Style.BodyBackgroundColor }} !important;
      color: {{ .Hermes.Style.BodyColor }};
      -webkit-text-size-adjust: none;
    }
    a {
      color: {{ .Hermes.Style.LinkColor }};
    }
    /* Layout ------------------------------ */
    .email-wrapper {
      width: 100%;
      margin: 0;
      padding: 0;
      background-color: {{ .Hermes.Style.EmailWrapperBackgroundColor }} !important;
    }
    .email-content {
      width: 100%;
      margin: 0;
      padding: 0;
    }
    /* Masthead ----------------------- */
    .email-masthead {
      padding: 25px 0;
      text-align: center;
    }
    .email-masthead_logo {
      max-width: 400px;
      border: 0;
    }
    .email-masthead_name {
      font-size: 16px;
      font-weight: bold;
      color: {{ .Hermes.Style.AppNameColor }};
      text-decoration: none;
      text-shadow: 0 0px 0 white;
    }
    .email-logo {
      max-height: 50px;
    }
    /* Body ------------------------------ */
    .email-body {
      width: 100%;
      margin: 0;
      padding: 0;
      border-top: 1px solid {{ .Hermes.Style.EmailBodyBorderTopColor }};
      border-bottom: 1px solid {{ .Hermes.Style.EmailBodyBorderBottomColor }};
      background-color: {{ .Hermes.Style.EmailBodyBackgroundColor }};
    }
    .email-body_inner {
      width: {{ .Hermes.Style.EmailBodyWidth }}px;
      margin: 0 auto;
      padding: 0;
    }
    .email-footer {
      width: 100% !important;
      margin: 0 auto;
      padding: 0;
      text-align: center;
      background-color: {{ .Hermes.Style.FooterBackgroundColor }};
    }
    .email-footer p {
      color: {{ .Hermes.Style.FooterColor }};
    }
    .body-action {
      width: 100%;
      margin: 30px auto;
      padding: 0;
      text-align: center;
    }
    .body-dictionary {
      width: 100%;
      overflow: hidden;
      margin: 20px auto 10px;
      padding: 0;
    }
    .body-dictionary dd {
      margin: 0 0 10px 0;
    }
    .body-dictionary dt {
      clear: both;
      color: #000;
      font-weight: bold;
    }
    .body-dictionary dd {
      margin-left: 0;
      margin-bottom: 10px;
    }
    .body-sub {
      margin-top: 25px;
      padding-top: 25px;
      border-top: 1px solid {{ .Hermes.Style.BodySubBorderTopColor }};
      table-layout: fixed;
    }
    .body-sub a {
      word-break: break-all;
    }
    .content-cell {
      padding: 35px;
    }
    .align-right {
      text-align: right;
    }
    /* Type ------------------------------ */
    h1 {
      margin-top: 0;
      color: {{ .Hermes.Style.H1Color }};
      font-size: 19px;
      font-weight: bold;
    }
    h2 {
      margin-top: 0;
      color: {{ .Hermes.Style.H2Color }};
      font-size: 16px;
      font-weight: bold;
    }
    h3 {
      margin-top: 0;
      color: {{ .Hermes.Style.H3Color }};
      font-size: 14px;
      font-weight: bold;
    }
    blockquote {
      margin: 25px 0;
      padding-left: 10px;
      border-left: 10px solid {{ .Hermes.Style.BlockquoteBorderLeftColor }};
    }
    blockquote p {
        font-size: 1.1rem;
        color: {{ .Hermes.Style.BlockquoteParagraphColor }};
    }
    blockquote cite {
        display: block;
        text-align: right;
        color: {{ .Hermes.Style.BlockquoteCiteColor }};
        font-size: 1.2rem;
    }
    cite {
      display: block;
      font-size: 0.925rem; 
    }
    cite:before {
      content: "\2014 \0020";
    }
    p {
      margin-top: 0;
      color: {{ .Hermes.Style.ParagraphColor }};
      font-size: 16px;
      line-height: 1.5em;
    }
    p.sub {
      font-size: 12px;
    }
    p.center {
      text-align: center;
    }
    table {
      width: 100%;
    }
    th {
      padding: 0px 5px;
      padding-bottom: 8px;
      border-bottom: 1px solid {{ .Hermes.Style.TableHeadingBorderBottomColor }};
    }
    th p {
      margin: 0;
      color: {{ .Hermes.Style.TableHeadingColor }};
      font-size: 12px;
    }
    td {
      padding: 10px 5px;
      color: {{ .Hermes.Style.TableDataColor }};
      font-size: 15px;
      line-height: 18px;
    }
    td.clear {
      padding: 0px 0px;
    }
    .content {
      align: center;
      padding: 0;
    }
    /* Data table ------------------------------ */
    .data-wrapper {
      width: 100%;
      margin: 0;
      padding: 35px 0;
    }
    .data-table {
      width: 100%;
      margin: 0;
    }
    .data-table th {
      text-align: left;
      padding: 0px 5px;
      padding-bottom: 8px;
      border-bottom: 1px solid {{ .Hermes.Style.DataTableTableHeadingBorderBottomColor }};
    }
    .data-table th p {
      margin: 0;
      color: {{ .Hermes.Style.DataTableTableHeadingColor }};
      font-size: 12px;
    }
    .data-table td {
      padding: 10px 5px;
      color: {{ .Hermes.Style.DataTableTableDataColor }};
      font-size: 15px;
      line-height: 18px;
    }
    /* Invite Code ------------------------------ */
    .invite-code {
      display: inline-block;
      padding-top: 20px;
      padding-right: 36px;
      padding-bottom: 16px;
      padding-left: 36px;
      border-radius: 3px;
      font-family: Consolas, monaco, monospace;
      font-size: 28px;
      text-align: center;
      letter-spacing: 8px;
      color: {{ .Hermes.Style.InviteCodeColor }};
      background-color: {{ .Hermes.Style.InviteCodeBackgroundColor }};
    }
    /* Buttons ------------------------------ */
    .button {
      display: inline-block;
      background-color: {{ .Hermes.Style.ActionButtonBackgroundColor }};
      border-radius: 3px;
      color: {{ .Hermes.Style.ActionButtonColor }} !important;
      font-size: 15px;
      line-height: 45px;
      text-align: center;
      text-decoration: none;
      -webkit-text-size-adjust: none;
      mso-hide: all;
    }
    /*Media Queries ------------------------------ */
    @media only screen and (max-width: 600px) {
      .email-body_inner,
      .email-footer {
        width: 100% !important;
      }
    }
    @media only screen and (max-width: 500px) {
      .button {
        width: 100% !important;
      }
    }
  </style>
</head>
<body dir="{{.Hermes.TextDirection}}">
  <table class="email-wrapper" width="100%" cellpadding="0" cellspacing="0">
    <tr>
      <td class="content">
        <table class="email-content" width="100%" cellpadding="0" cellspacing="0">
          <!-- Logo -->
          <tr>
            <td class="email-masthead">
                {{ if .Hermes.Product.Link }}<a class="email-masthead_name" href="{{.Hermes.Product.Link}}" target="_blank">{{ end }}
                {{ if .Hermes.Product.Logo }}
                  <img src="{{.Hermes.Product.Logo | url }}" class="email-logo" />
                {{ else }}
                  {{ .Hermes.Product.Name }}
                {{ end }}
                {{ if .Hermes.Product.Link }}</a>{{ end }}
            </td>
          </tr>

          <!-- Email Body -->
          <tr>
            <td class="email-body" width="100%">
              <table class="email-body_inner" align="center" width="{{ .Hermes.Style.EmailBodyWidth }}" cellpadding="0" cellspacing="0">
                <!-- Body content -->
                <tr>
                  <td class="content-cell">
                    <h1>{{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }} {{ .Email.Body.Name }},{{ end }}</h1>
                    {{ with .Email.Body.Intros }}
                        {{ if gt (len .) 0 }}
                          {{ range $line := . }}
                            <p>{{ $line }}</p>
                          {{ end }}
                        {{ end }}
                    {{ end }}
                    {{ if (ne .Email.Body.IntroMarkdown "") }}
                      {{ .Email.Body.IntroMarkdown.ToHTML }}
                    {{ end }}
                    {{ if (ne .Email.Body.FreeMarkdown "") }}
                      {{ .Email.Body.FreeMarkdown.ToHTML }}
                    {{ else }}

                      {{ with .Email.Body.Dictionary }} 
                        {{ if gt (len .) 0 }}
                          <dl class="body-dictionary">
                            {{ range $entry := . }}
                              <dt>{{ $entry.Key }}:</dt>
                              <dd>{{ $entry.Value }}</dd>
                            {{ end }}
                          </dl>
                        {{ end }}
                      {{ end }}

                      <!-- Table -->
                      {{ with .Email.Body.Table }}
                        {{ $data := .Data }}
                        {{ $columns := .Columns }}
                        {{ if gt (len $data) 0 }}
                          <table class="data-wrapper" width="100%" cellpadding="0" cellspacing="0">
                            <tr>
                              <td colspan="2">
                                <table class="data-table" width="100%" cellpadding="0" cellspacing="0">
                                  <tr>
                                    {{ $col := index $data 0 }}
                                    {{ range $entry := $col }}
                                      <th
                                        {{ with $columns }}
                                          {{ $width := index .CustomWidth $entry.Key }}
                                          {{ with $width }}
                                            width="{{ . }}"
                                          {{ end }}
                                          {{ $align := index .CustomAlignment $entry.Key }}
                                          {{ with $align }}
                                            style="text-align:{{ . }}"
                                          {{ end }}
                                        {{ end }}
                                      >
                                        <p>{{ $entry.Key }}</p>
                                      </th>
                                    {{ end }}
                                  </tr>
                                  {{ range $row := $data }}
                                    <tr>
                                      {{ range $cell := $row }}
                                        <td
                                          {{ with $columns }}
                                            {{ $align := index .CustomAlignment $cell.Key }}
                                            {{ with $align }}
                                              style="text-align:{{ . }}"
                                            {{ end }}
                                          {{ end }}
                                        >
                                          {{ $cell.Value }}
                                        </td>
                                      {{ end }}
                                    </tr>
                                  {{ end }}
                                </table>
                              </td>
                            </tr>
                          </table>
                        {{ end }}
                      {{ end }}

                      <!-- Action -->
                      {{ with .Email.Body.Actions }}
                        {{ if gt (len .) 0 }}
                          {{ range $action := . }}
                            {{ if (ne $action.InstructionsMarkdown "") }}
                              {{ $action.InstructionsMarkdown.ToHTML }}
                            {{ else }}
                              <p>{{ $action.Instructions }}</p>
                            {{ end }}
                            {{ $length := len $action.Button.Text }}
                            {{ $width := add (mul $length 9) 100 }}
                            {{if (lt $width 200)}}{{$width = 200}}{{else if (gt $width $.Hermes.Style.EmailBodyWidth )}}{{$width = $.Hermes.Style.EmailBodyWidth }}{{else}}{{end}}
                              {{safe "<!--[if mso]>" }}
                              {{ if $action.Button.Text }}
                                <div style="margin: 30px auto;v-text-anchor:middle;text-align:center">
                                  <v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" 
                                    xmlns:w="urn:schemas-microsoft-com:office:word" 
                                    href="{{ $action.Button.Link }}" 
                                    style="height:45px;v-text-anchor:middle;width:{{$width}}px;background-color:{{ if $action.Button.Color }}{{ $action.Button.Color }}{{ else }}{{ $.Hermes.Style.ActionButtonBackgroundColor }}{{ end }};"
                                    arcsize="10%" 
                                    {{ if $action.Button.Color }}strokecolor="{{ $action.Button.Color }}" fillcolor="{{ $action.Button.Color }}"{{ else }}strokecolor="{{ $.Hermes.Style.ActionButtonBackgroundColor }}" fillcolor="{{ $.Hermes.Style.ActionButtonBackgroundColor }}"{{ end }}
                                    >
                                    <w:anchorlock/>
                                    <center style="color: {{ if $action.Button.TextColor }}{{ $action.Button.TextColor }}{{else}}{{ $.Hermes.Style.ActionButtonColor }}{{ end }};font-size: 15px;text-align: center;font-family:sans-serif;font-weight:bold;">
                                      {{ $action.Button.Text }}
                                    </center>
                                  </v:roundrect>
                                </div>
                              {{ end }}
                              {{ if $action.InviteCode }}
                                <div style="margin-top:30px;margin-bottom:30px">
                                  <table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0">
                                    <tr>
                                      <td align="center">
                                        <table align="center" cellpadding="0" cellspacing="0" style="padding:0;text-align:center">
                                          <tr>
                                            <td style="display:inline-block;border-radius:3px;font-family:Consolas, monaco, monospace;font-size:28px;text-align:center;letter-spacing:8px;color:{{ $.Hermes.Style.InviteCodeColor }};background-color:{{ $.Hermes.Style.InviteCodeBackgroundColor }};padding:20px">
                                              {{ $action.InviteCode }}
                                            </td>
                                          </tr>
                                        </table>
                                      </td>
                                    </tr>
                                  </table>
                                </div>
                              {{ end }}   
                              {{safe "<![endif]-->" }}
                              {{safe "<!--[if !mso]><!-- -->"}}
                              <table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0">
                                <tr>
                                  <td align="center">
                                    <div>
                                      {{ if $action.Button.Text }}
                                        <a href="{{ $action.Button.Link }}" class="button" style="{{ with $action.Button.Color }}background-color: {{ . }};{{ end }} {{ with $action.Button.TextColor }}color: {{ . }};{{ end }} width: {{$width}}px;" target="_blank">
                                          {{ $action.Button.Text }}
                                        </a>
                                      {{end}}
                                      {{ if $action.InviteCode }}
                                        <span class="invite-code">{{ $action.InviteCode }}</span>
                                      {{end}}
                                    </div>
                                  </td>
                                </tr>
                              </table>
                              {{safe "<![endif]-->" }}
                          {{ end }}
                        {{ end }}
                      {{ end }}

                    {{ end }}
                    {{ with .Email.Body.Outros }} 
                      {{ if gt (len .) 0 }}
                        {{ range $line := . }}
                          <p>{{ $line }}</p>
                        {{ end }}
                      {{ end }}
                    {{ end }}

                    {{ if (ne .Email.Body.OutroMarkdown "") }}
                      {{ .Email.Body.OutroMarkdown.ToHTML }}
                    {{ end }}

                    <p>
                      {{ if .Email.Body.Signature }}
                        {{.Email.Body.Signature}},
                        <br />
                        {{.Hermes.Product.Name}}
                      {{ end }}
                    </p>

                    {{ if (eq .Email.Body.FreeMarkdown "") }}
                      {{ with .Email.Body.Actions }} 
                        <table class="body-sub">
                          <tbody>
                              {{ range $action := . }}
                                {{if $action.Button.Text}}
                                <tr>
                                  <td>
                                    <p class="sub">{{$.Hermes.Product.TroubleText | replace "{ACTION}" $action.Button.Text}}</p>
                                    <p class="sub"><a href="{{ $action.Button.Link }}">{{ $action.Button.Link }}</a></p>
                                  </td>
                                </tr>
                                {{ end }}
                              {{ end }}
                          </tbody>
                        </table>
                      {{ end }}
                    {{ end }}
                  </td>
                </tr>
              </table>
            </td>
          </tr>
          <tr>
            <td class="clear">
              <table class="email-footer" align="center" width="{{ .Hermes.Style.EmailBodyWidth }}" cellpadding="0" cellspacing="0">
                <tr>
                  <td class="content-cell">
                    <p class="sub center">
                      {{.Hermes.Product.Copyright}}
                    </p>
                  </td>
                </tr>
              </table>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</body>
</html>
`
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *Custom) PlainTextTemplate() string {
	return `<h2>{{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }} {{ .Email.Body.Name }},{{ end }}</h2>
{{ with .Email.Body.Intros }}
  {{ range $line := . }}
    <p>{{ $line }}</p>
  {{ end }}
{{ end }}
{{ if (ne .Email.Body.IntroMarkdown "") }}
  {{ .Email.Body.IntroMarkdown.ToHTML }}
{{ end }}
{{ if (ne .Email.Body.FreeMarkdown "") }}
  {{ .Email.Body.FreeMarkdown.ToHTML }}
{{ else }}
  {{ with .Email.Body.Dictionary }}
    <ul>
    {{ range $entry := . }}
      <li>{{ $entry.Key }}: {{ $entry.Value }}</li>
    {{ end }}
    </ul>
  {{ end }}
  {{ with .Email.Body.Table }}
    {{ $data := .Data }}
    {{ $columns := .Columns }}
    {{ if gt (len $data) 0 }}
      <table class="data-table" width="100%" cellpadding="0" cellspacing="0">
        <tr>
          {{ $col := index $data 0 }}
          {{ range $entry := $col }}
            <th>{{ $entry.Key }} </th>
          {{ end }}
        </tr>
        {{ range $row := $data }}
          <tr>
            {{ range $cell := $row }}
              <td>
                {{ $cell.Value }}
              </td>
            {{ end }}
          </tr>
        {{ end }}
      </table>
    {{ end }}
  {{ end }}
  {{ with .Email.Body.Actions }} 
    {{ range $action := . }}
      <p>
        {{ if (ne $action.InstructionsMarkdown "") }}
          {{ $action.InstructionsMarkdown.ToHTML }}
        {{ else }}
          {{ $action.Instructions }}
        {{ end }}
        {{ if $action.InviteCode }}
          {{ $action.InviteCode }}
        {{ end }}
        {{ if $action.Button.Link }}
          {{ $action.Button.Link }}
        {{ end }}
      </p> 
    {{ end }}
  {{ end }}
{{ end }}
{{ with .Email.Body.Outros }} 
  {{ range $line := . }}
    <p>{{ $line }}<p>
  {{ end }}
{{ end }}
{{ if (ne .Email.Body.OutroMarkdown "") }}
  {{ .Email.Body.OutroMarkdown.ToHTML }}
{{ end }}
<p>
{{ if .Email.Body.Signature }}
  {{.Email.Body.Signature}},
  <br />
  {{.Hermes.Product.Name}} - {{.Hermes.Product.Link}}
{{ end }}
</p>
<p>{{.Hermes.Product.Copyright}}</p>
`
}
