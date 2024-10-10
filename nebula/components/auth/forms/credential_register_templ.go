// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package forms

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func RegisterCredentialForm() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"border rounded-lg shadow-sm bg-card text-zinc-900\"><div class=\"flex flex-col space-y-1.5 p-6\"></div><div class=\"p-6 pt-0 space-y-2\"><div class=\"space-y-1\"><label class=\"text-xs font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"name\">Name</label><input type=\"text\" id=\"name\" placeholder=\"Adam Wathan\" class=\"flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md peer border-zinc-300 ring-offset-background placeholder:text-zinc-400 focus:border-zinc-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50\"></div><div class=\"space-y-1\"><label class=\"text-xs font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"username\">Handle</label><input type=\"text\" id=\"handle\" placeholder=\"angelo.snr\" class=\"flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md peer border-zinc-300 ring-offset-background placeholder:text-zinc-400 focus:border-zinc-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
