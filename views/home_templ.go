// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Homepage() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.Raw("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"https://unpkg.com/htmx.org@1.9.12\" type=\"text/javascript\"></script><script src=\"https://cdn.tailwindcss.com\"></script><title>WaMeir - Reach out to WhatsApp users without saving numbers</title></head><body class=\"bg-gray-100 text-gray-800\"><div class=\"container mx-auto p-6\"><h2 class=\"text-3xl font-bold text-center mb-6 text-green-900\">WaMeir 😀👌</h2><ul class=\"list-disc list-inside mb-6\"><li class=\"mb-2\">For Nigeria Numbers: Enter your numbers in this format 080XXXXXXXXX</li><li class=\"mb-2\">For USA Numbers: Enter your numbers in this format +1XXXXXXXXXX</li><li class=\"mb-2\">For UK Numbers: Enter your numbers in this format +44XXXXXXXXXX</li></ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = checkWalink().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func checkWalink() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"flex flex-col gap-6\" action=\"/lists\" method=\"post\" hx-swap=\"transition:true\"><div class=\"mb-4\"><label for=\"Country\" class=\"block text-gray-700 text-sm font-bold mb-2\">Choose Your Country:</label> <select id=\"cnid\" name=\"cid\" class=\"shadow appearance-none border border-green-700 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline\"><option value=\"https://wa.me/+234\">Nigeria</option> <option value=\"https://wa.me/+1\">United States</option> <option value=\"https://wa.me/+44\">United Kingdom</option></select></div><div class=\"mb-4\"><label for=\"numbs\" class=\"block text-gray-700 text-sm font-bold mb-2\">Paste those numbers:</label> <textarea id=\"numbs\" name=\"numbs\" rows=\"10\" cols=\"50\" placeholder=\"Paste those numbers\" class=\"shadow appearance-none border border-green-700 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline\"></textarea></div><div class=\"flex items-center justify-between\"><button type=\"submit\" class=\"bg-green-700 hover:bg-green-800 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline\">Submit</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
