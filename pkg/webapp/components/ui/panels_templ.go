// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Breadcrumbs() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"flex justify-between px-3.5 py-1 rounded-md pb-3\"><ol class=\"inline-flex items-center mb-3 space-x-1 text-xs text-neutral-500 [&amp;_.active-breadcrumb]:text-neutral-600 [&amp;_.active-breadcrumb]:font-medium sm:mb-0\"><li class=\"flex items-center h-full\"><a href=\"#_\" class=\"py-1 hover:text-neutral-900\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = breadcrumbIcon().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li><svg class=\"w-5 h-5 text-gray-400/70\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><g fill=\"none\" stroke=\"none\"><path d=\"M10 8.013l4 4-4 4\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = breadcrumbItem("Policy", false).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg class=\"w-5 h-5 text-gray-400/70\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><g fill=\"none\" stroke=\"none\"><path d=\"M10 8.013l4 4-4 4\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = breadcrumbItem("About You", true).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg class=\"w-5 h-5 text-gray-400/70\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><g fill=\"none\" stroke=\"none\"><path d=\"M10 8.013l4 4-4 4\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = breadcrumbItem("Generate", false).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ol></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func breadcrumbItem(title string, active bool) templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if active {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a class=\"inline-flex items-center py-1 font-normal rounded cursor-default active-breadcrumb focus:outline-none\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/webapp/components/ui/panels.templ`, Line: 23, Col: 126}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a href=\"#_\" class=\"inline-flex items-center py-1 font-normal hover:text-neutral-900 focus:outline-none\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/webapp/components/ui/panels.templ`, Line: 25, Col: 118}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func breadcrumbIcon() templ.Component {
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
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg class=\"w-4 h-4\" viewBox=\"0 0 164 164\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M71.8077 133.231C74.5054 135.928 78.1636 137.443 81.978 137.443C85.7924 137.443 89.4506 135.928 92.1483 133.231L133.219 92.1638C135.909 89.4654 137.42 85.8102 137.42 81.9998C137.42 78.1895 135.909 74.5345 133.219 71.8361L112.886 51.5272L131.665 32.7499L152.031 53.1143C159.696 60.7963 164 71.2046 164 82.0559C164 92.9072 159.696 103.315 152.031 110.997L110.95 152.065C107.154 155.869 102.642 158.883 97.6739 160.931C92.7059 162.98 87.3809 164.023 82.0071 164L82.0052 164C76.622 164.019 71.2886 162.969 66.3145 160.91C61.3405 158.852 56.8247 155.826 53.0294 152.009L53.0289 152.008L48.7187 147.699L67.4974 128.921L71.8077 133.231Z\" fill=\"currentColor\"></path> <path d=\"M110.95 11.9912L115.26 16.3011L96.481 35.0785L92.1707 30.7685C89.4731 28.072 85.8148 26.5572 82.0004 26.5572C78.186 26.5572 74.5277 28.072 71.8301 30.7685L30.7597 71.8359C29.4247 73.1706 28.3658 74.7552 27.6433 76.4991C26.9208 78.2431 26.549 80.1122 26.549 81.9999C26.549 83.8876 26.9208 85.7567 27.6433 87.5007C28.3658 89.2446 29.4247 90.8292 30.7597 92.1639L51.1256 112.528L32.3138 131.306L11.9923 110.941C8.19043 107.141 5.17433 102.629 3.1167 97.6635C1.05907 92.6976 0 87.3751 0 81.9999C0 76.6247 1.05907 71.3022 3.1167 66.3363C5.17433 61.3705 8.19021 56.8587 11.9921 53.0586L53.0625 11.9912C56.8629 8.18964 61.3751 5.17395 66.3413 3.11647C71.3075 1.05899 76.6304 0 82.006 0C87.3816 0 92.7045 1.05899 97.6707 3.11647C102.637 5.17395 107.149 8.18964 110.95 11.9912Z\" fill=\"currentColor\"></path> <path d=\"M55.603 76.6744L76.6993 55.5798C79.6327 52.6465 84.3888 52.6465 87.3223 55.5797L108.419 76.6744C111.352 79.6077 111.352 84.3634 108.419 87.2966L87.3223 108.391C84.3888 111.325 79.6327 111.325 76.6993 108.391L55.603 87.2966C52.6696 84.3634 52.6696 79.6077 55.603 76.6744Z\" fill=\"currentColor\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
