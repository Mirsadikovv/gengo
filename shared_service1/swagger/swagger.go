package swagger

import (
	"bytes"
	"embed"
	"encoding/json"
	"html/template"
	"io"
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/Mirsadikovv/gengo/shared_service/sharedutil"
	"github.com/labstack/echo/v4"
)

//go:embed swagger-ui
var swagger embed.FS

func NewSwaggerHandler(docs fs.FS, prefix ...string) echo.HandlerFunc {
	swaggerfs := sharedutil.MustValue(fs.Sub(swagger, "swagger-ui"))

	urls := []map[string]any{}

	var url string = "/swagger/"
	{
		if len(prefix) > 0 {
			url = prefix[0]
		}
	}

	fs.WalkDir(docs, ".", func(sysPath string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		paths := strings.Split(sysPath, "/")

		urls = append(urls, map[string]any{
			"name": paths[len(paths)-2],
			"url":  path.Join(url, sysPath),
		})

		return nil
	})

	data, err := swagger.ReadFile("swagger-ui/index.html")
	{
		if err != nil {
			panic(err)
		}
	}

	var indexData string
	{
		var buff bytes.Buffer
		template.Must(template.New("").
			Funcs(template.FuncMap{
				"json": func(v any) string {
					b, _ := json.Marshal(v)
					return string(b)
				},
			}).
			Parse(string(data))).
			Execute(&buff, urls)
		indexData = buff.String()
	}

	mimeType := func(fileName string) string {
		ext := strings.ToLower(filepath.Ext(fileName))
		switch ext {
		case ".css":
			return "text/css"
		case ".js":
			return "application/javascript"
		case ".json":
			return "application/json"
		default:
			return "application/octet-stream"
		}
	}

	return func(c echo.Context) error {
		path := c.Param("*")
		{
			if path == "" {
				return c.HTML(200, indexData)
			}
		}

		c.Response().Header().Set("Content-Type", mimeType(path))

		if strings.HasSuffix(path, ".json") {

			f, err := docs.Open(path)
			{
				if err != nil {
					return c.String(404, err.Error())
				}
			}

			if _, err := io.Copy(c.Response(), f); err != nil {
				return c.String(404, err.Error())
			}

			return nil
		}

		f, err := swaggerfs.Open(path)
		{
			if err != nil {
				return c.String(404, err.Error())
			}
		}

		if _, err := io.Copy(c.Response(), f); err != nil {
			return c.String(404, err.Error())
		}

		return nil
	}
}
