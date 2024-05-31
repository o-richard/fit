// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func changeTheme() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_changeTheme_e9c7`,
		Function: `function __templ_changeTheme_e9c7(){document.documentElement.classList.toggle('dark')
	localStorage.fitTheme = (document.documentElement.classList.contains('dark')) ? 'dark' : 'light'
}`,
		Call:       templ.SafeScript(`__templ_changeTheme_e9c7`),
		CallInline: templ.SafeScriptInline(`__templ_changeTheme_e9c7`),
	}
}

func base(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"robots\" content=\"noindex, nofollow\"><link rel=\"stylesheet\" href=\"/css/styles.css\"><link rel=\"shortcut icon\" href=\"/favicon.ico\" type=\"image/x-icon\"><link rel=\"apple-touch-icon\" href=\"/img/logo.jpg\"><title>Fit  ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if title != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("| ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/base.templ`, Line: 21, Col: 14}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><script>\n\t\t\t\tif (localStorage.fitTheme === 'dark' || (!('fitTheme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {\n\t\t\t\t\tdocument.documentElement.classList.add('dark')\n\t\t\t\t\tlocalStorage.fitTheme = 'dark'\n\t\t\t\t} else {\n\t\t\t\t\tdocument.documentElement.classList.remove('dark')\n\t\t\t\t\tlocalStorage.fitTheme = 'light'\n\t\t\t\t}\n\t\t\t</script><script src=\"/js/htmx.min.js\"></script></head><body class=\"bg-white text-black dark:bg-black dark:text-white px-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, changeTheme())
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button onClick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.ComponentScript = changeTheme()
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-red-600 text-lg absolute top-2\">Change Light/Dark Mode</button><main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
