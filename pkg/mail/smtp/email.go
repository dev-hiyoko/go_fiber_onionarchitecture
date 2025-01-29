package smtp

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var (
	templateDirectory   = "./pkg/mail/templates/"
	commonDirectoryName = "common/"
)

type Email struct {
	From             string
	FromName         string
	To               []string
	Cc               []string
	Bcc              []string
	Subject          string
	TemplateFileName string
	Data             map[string]string
	LangData         map[string]string
}

func NewEmail(
	from string,
	fromName string,
) *Email {
	return &Email{
		From:     from,
		FromName: fromName,
	}
}

func SetTemplateDirectory(dir string) {
	templateDirectory = dir
}

func SetCommonDirectoryName(dirName string) {
	commonDirectoryName = dirName
}

// SetData Add all data in map.
func (e *Email) SetData(data map[string]string) {
	for k, v := range data {
		e.Data[k] = v
	}
}

// SetAssetsPngImages sets the image images in the directory to the image structure.
// ※ only one directory hierarchy, only png supported
func (e *Email) SetAssetsPngImages(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	if e.Data == nil {
		e.Data = make(map[string]string)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext != ".png" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return err
		}

		enc := base64.StdEncoding.EncodeToString(data)

		name := strings.TrimSuffix(entry.Name(), ext)
		e.Data["Assets/Images/"+name] = enc
	}

	return nil
}

func (e *Email) buildMsgFromTemplate() (string, error) {
	commonDir := filepath.Join(templateDirectory, commonDirectoryName)
	templateFiles, err := filepath.Glob(commonDir + "/*.tmpl")
	if err != nil {
		return "", err
	}

	t, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return "", err
	}

	t, err = t.ParseFiles(filepath.Join(templateDirectory, e.TemplateFileName))
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.ExecuteTemplate(buf, e.TemplateFileName, e.Data); err != nil {
		return "", err
	}

	header := make(map[string]string)
	header["From"] = fmt.Sprintf("%s <%s>", e.FromName, e.From)
	header["To"] = strings.Join(e.To, ",")
	header["Cc"] = strings.Join(e.Cc, ",")
	header["Subject"] = e.Subject
	header["Content-Type"] = `text/html; charset="UTF-8"`

	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	msg += "\r\n" + buf.String()
	return msg, nil
}

func (e *Email) buildData() {
	if e.Data == nil {
		e.Data = make(map[string]string)
		return
	}

	for k, v := range e.Data {
		if strings.HasPrefix(v, "{{") && strings.HasSuffix(v, "}}") {
			key := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(v, "{{"), "}}"))
			key = strings.TrimPrefix(key, ".")
			e.Data[k] = e.LangData[key]
		}
	}

	for k, v := range e.LangData {
		e.Data[k] = v
	}
}

func (e *Email) createRecipientList() []string {
	recipients := append(e.To, e.Cc...)
	recipients = append(recipients, e.Bcc...)
	return recipients
}
