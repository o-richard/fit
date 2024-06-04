// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"

func insertHealthEntry() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_insertHealthEntry_2444`,
		Function: `function __templ_insertHealthEntry_2444(){const gallery = document.getElementById("gallery")
	const noImages = document.getElementById("no-images")
	const imageTemplate = document.getElementById("image-template")
	const dropzoneFileInput = document.getElementById("dropzone-file")
	
	const formBtn = document.getElementById("submit-button")
	const formElement = document.getElementById("new-entry-form")
	const formContent = document.getElementById("form-content")
	const formEndedAt = document.getElementById("form-ended-at")
	const formTimezone = document.getElementById("formTimezone")
	const formStartedAt = document.getElementById("form-started-at")
	
	let selectedImages = {}
	function addFile(target, file) {
		const oneMb = 1024 * 1024
		if ((!(file.type === "image/png" || file.type === "image/jpeg")) || (file.size > (2 * oneMb))) {
			return
		}
		
		const imageObjectURL = URL.createObjectURL(file)
		const imageClone = imageTemplate.content.cloneNode(true)
		let humanReadableSize;
		if (file.size > 1024) {
			humanReadableSize = file.size > oneMb ? Math.round(file.size / oneMb) + "mb" : Math.round(file.size / 1024) + "kb"
		} else {
			humanReadableSize = file.size + "b"
		}

		imageClone.querySelector("li").id = imageObjectURL
		imageClone.querySelector("h1").textContent = file.name
		imageClone.querySelector(".delete").dataset.target = imageObjectURL
		imageClone.querySelector(".size").textContent = humanReadableSize
		Object.assign(imageClone.querySelector("img"), {src: imageObjectURL, alt: file.name})

		noImages.classList.add("hidden")
		target.prepend(imageClone)
		selectedImages[imageObjectURL] = file
	}
	gallery.addEventListener('click', ({ target }) => {
		if (target.classList.contains("delete")) {
			const imageObjectURL = target.dataset.target
			document.getElementById(imageObjectURL).remove()
			gallery.children.length === 1 && noImages.classList.remove("hidden")
			delete selectedImages[imageObjectURL]
		}    
	})
	dropzoneFileInput.addEventListener('change', (e) => {
		for (const file of e.target.files) {
			addFile(gallery, file)
		}
		const dataTransfer = new DataTransfer()
		for (const [_, value] of Object.entries(selectedImages)) {
			dataTransfer.items.add(value)
		}
		dropzoneFileInput.files = dataTransfer.files
	})

	formTimezone.value = Intl.DateTimeFormat().resolvedOptions().timeZone
	formContent.addEventListener("change", () => {
		if (formContent.value.trim() === "") {
			formContent.setCustomValidity("Please fill out this field.")
		} else {
			formContent.setCustomValidity("")
		}
	})
	formStartedAt.addEventListener("change", () => {
		if (formEndedAt.value === "") {
			formEndedAt.value = formStartedAt.value
		}
	})
	formEndedAt.addEventListener("change", () => {
		if (formStartedAt.value !== "") {
			const startedAtTimestamp = new Date(formStartedAt.value)
			const endedAtTimestamp = new Date(formEndedAt.value)
			if (endedAtTimestamp.getTime() < startedAtTimestamp.getTime()) {
				formEndedAt.setCustomValidity("The end timestamp should be greater than or equal to the start timestamp.")
			} else {
				formEndedAt.setCustomValidity("")
			}
		}
	})
}`,
		Call:       templ.SafeScript(`__templ_insertHealthEntry_2444`),
		CallInline: templ.SafeScriptInline(`__templ_insertHealthEntry_2444`),
	}
}

func InsertHealthEntry() templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"flex flex-col items-center\"><h1 class=\"text-2xl font-bold py-4\">New Entry</h1><form enctype=\"multipart/form-data\" class=\"w-4/5 lg:w-2/5 pb-4\" id=\"new-entry-form\" hx-post=\"/entry/new\" hx-target=\"#toastr-notifications\" hx-swap=\"innerHTML show:#main-content:top\" hx-on::after-request=\"this.reset()\"><div class=\"flex flex-col gap-2\"><label for=\"entryType\" class=\"after:content-[&#39;*&#39;] after:ml-0.5 after:text-red-500 text-base font-bold\">Type</label> <select name=\"entryType\" id=\"entryType\" class=\"rounded-lg p-4 font-bold bg-gray-100 text-gray-700 dark:bg-gray-700 dark:text-gray-100 focus:outline-none\"><option value=\"health\" selected>Health</option> <option value=\"activity\">Activity</option> <option value=\"sleep\">Sleep</option> <option value=\"nutrition\">Nutrition</option></select></div><div class=\"flex flex-col gap-2 py-2\"><label for=\"title\" class=\"text-base font-bold\">Title</label> <input type=\"text\" name=\"title\" id=\"title\" class=\"rounded-lg p-4 font-bold bg-gray-100 text-gray-700 dark:bg-gray-700 dark:text-gray-100 focus:outline-none\"></div><div class=\"flex flex-col gap-2 py-2\"><label for=\"form-content\" class=\"after:content-[&#39;*&#39;] after:ml-0.5 after:text-red-500 text-base font-bold\">Content</label> <textarea id=\"form-content\" name=\"content\" rows=\"4\" class=\"rounded-lg p-4 font-bold bg-gray-100 text-gray-700 dark:bg-gray-700 dark:text-gray-100 focus:outline-none\" required></textarea></div><input type=\"hidden\" name=\"timezone\" value=\"\" id=\"formTimezone\"><div class=\"flex flex-col gap-2 py-2\"><label for=\"dropzone-file\" class=\"flex items-center justify-center h-40 border-2 border-gray-300 dark:border-gray-600 dark:hover:border-gray-500 border-dashed rounded-lg cursor-pointer bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-bray-800\"><div class=\"flex flex-col items-center pt-5 pb-6\"><svg class=\"w-8 h-8 mb-2 text-gray-700 dark:text-gray-200\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 20 16\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 13h3a3 3 0 0 0 0-6h-.025A5.56 5.56 0 0 0 16 6.5 5.5 5.5 0 0 0 5.207 5.021C5.137 5.017 5.071 5 5 5a4 4 0 0 0 0 8h2.167M10 15V6m0 0L8 8m2-2 2 2\"></path></svg><p class=\"mb-2 text-base font-semibold text-gray-700 dark:text-gray-200\">Click to upload</p><p class=\"text-xs text-gray-700 dark:text-gray-200\">JPG or PNG (MAX. 2MB)</p></div><input id=\"dropzone-file\" type=\"file\" name=\"images\" multiple class=\"hidden\" accept=\"image/png, image/jpeg\"></label><ul id=\"gallery\" class=\"flex flex-1 flex-wrap -m-1\"><li id=\"no-images\" class=\"h-44 w-full flex flex-col justify-center \"><svg class=\"mx-auto\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" version=\"1.1\" width=\"100\" height=\"100\" viewBox=\"0 0 256 256\" xml:space=\"preserve\"><defs></defs> <g style=\"stroke: none; stroke-width: 0; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: none; fill-rule: nonzero; opacity: 1;\" transform=\"translate(22.611673151750978 22.61167315175095) scale(2.33 2.33)\"><path d=\"M 89 20.938 c -0.553 0 -1 0.448 -1 1 v 46.125 c 0 2.422 -1.135 4.581 -2.898 5.983 L 62.328 50.71 c -0.37 -0.379 -0.973 -0.404 -1.372 -0.057 L 45.058 64.479 l -2.862 -2.942 c -0.385 -0.396 -1.019 -0.405 -1.414 -0.02 c -0.396 0.385 -0.405 1.019 -0.02 1.414 l 3.521 3.62 c 0.37 0.38 0.972 0.405 1.373 0.058 l 15.899 -13.826 l 21.783 22.32 c -0.918 0.391 -1.928 0.608 -2.987 0.608 H 24.7 c -0.552 0 -1 0.447 -1 1 s 0.448 1 1 1 h 55.651 c 5.32 0 9.648 -4.328 9.648 -9.647 V 21.938 C 90 21.386 89.553 20.938 89 20.938 z\" style=\"stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(96,165,250); fill-rule: nonzero; opacity: 1;\" transform=\" matrix(1 0 0 1 0 0) \" stroke-linecap=\"round\"></path> <path d=\"M 89.744 4.864 c -0.369 -0.411 -1.002 -0.444 -1.412 -0.077 l -8.363 7.502 H 9.648 C 4.328 12.29 0 16.618 0 21.938 v 46.125 c 0 4.528 3.141 8.328 7.356 9.361 l -7.024 6.3 c -0.411 0.368 -0.445 1.001 -0.077 1.412 c 0.198 0.22 0.471 0.332 0.745 0.332 c 0.238 0 0.476 -0.084 0.667 -0.256 l 88 -78.935 C 90.079 5.908 90.113 5.275 89.744 4.864 z M 9.648 14.29 h 68.091 L 34.215 53.33 L 23.428 42.239 c -0.374 -0.385 -0.985 -0.404 -1.385 -0.046 L 2 60.201 V 21.938 C 2 17.721 5.431 14.29 9.648 14.29 z M 2 68.063 v -5.172 l 20.665 -18.568 l 10.061 10.345 L 9.286 75.692 C 5.238 75.501 2 72.157 2 68.063 z\" style=\"stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(96,165,250); fill-rule: nonzero; opacity: 1;\" transform=\" matrix(1 0 0 1 0 0) \" stroke-linecap=\"round\"></path> <path d=\"M 32.607 35.608 c -4.044 0 -7.335 -3.291 -7.335 -7.335 s 3.291 -7.335 7.335 -7.335 s 7.335 3.291 7.335 7.335 S 36.652 35.608 32.607 35.608 z M 32.607 22.938 c -2.942 0 -5.335 2.393 -5.335 5.335 s 2.393 5.335 5.335 5.335 s 5.335 -2.393 5.335 -5.335 S 35.549 22.938 32.607 22.938 z\" style=\"stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(96,165,250); fill-rule: nonzero; opacity: 1;\" transform=\" matrix(1 0 0 1 0 0) \" stroke-linecap=\"round\"></path></g></svg> <span class=\"text-base font-bold text-center text-gray-700 dark:text-gray-100\">No images selected</span></li></ul><template id=\"image-template\"><li class=\"block p-1 w-1/2 sm:w-1/3 md:w-1/4 lg:w-1/6 xl:w-1/8 h-24\"><article tabindex=\"0\" class=\"group w-full h-full rounded-md focus:outline-none focus:shadow-outline bg-gray-100 cursor-pointer relative shadow-sm text-transparent hover:text-white\"><img alt=\"upload preview\" class=\"w-full h-full sticky object-cover rounded-md bg-fixed\"><section class=\"flex flex-col rounded-md text-xs break-words w-full h-full z-20 absolute top-0 py-2 px-3 group-hover:bg-[#05050566]\"><h1 class=\"flex-1\"></h1><div class=\"flex\"><p class=\"p-1 size text-xs\"></p><button class=\"delete ml-auto focus:outline-none hover:bg-gray-300 p-1 rounded-md group-hover:hover:bg-[#05050573]\" aria-label=\"Delete\"><svg class=\"pointer-events-none fill-current w-4 h-4 ml-auto\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\"><path class=\"pointer-events-none\" d=\"M3 6l3 18h12l3-18h-18zm19-4v2h-20v-2h5.711c.9 0 1.631-1.099 1.631-2h5.316c0 .901.73 2 1.631 2h5.711z\"></path></svg></button></div></section></article></li></template></div><div class=\"grid grid-cols-1 md:grid-cols-2 gap-4 py-2\"><div class=\"col-span-1 flex flex-col\"><label for=\"form-started-at\" class=\"after:content-[&#39;*&#39;] after:ml-0.5 after:text-red-500 text-base font-bold\">Started At</label> <input type=\"datetime-local\" name=\"startedAt\" id=\"form-started-at\" class=\"rounded-lg p-4 font-bold bg-gray-100 text-gray-700 dark:bg-gray-700 dark:text-gray-100 focus:outline-none\" required></div><div class=\"col-span-1 flex flex-col\"><label for=\"form-ended-at\" class=\"after:content-[&#39;*&#39;] after:ml-0.5 after:text-red-500 text-base font-bold\">Ended At</label> <input type=\"datetime-local\" name=\"endedAt\" id=\"form-ended-at\" class=\"rounded-lg p-4 font-bold bg-gray-100 text-gray-700 dark:bg-gray-700 dark:text-gray-100 focus:outline-none\" required></div></div><button type=\"submit\" class=\"mt-2 py-2 w-full bg-blue-400 hover:bg-blue-600 text-base text-white font-bold rounded-2xl transition duration-200\" id=\"submit-button\">Submit</button></form></main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = base("New Health Entry", insertHealthEntry()).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func individualEntryCard(bgColor, content, urlparams string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full h-52\"><a hx-trigger=\"click\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(string(templ.URL(fmt.Sprintf("/entry?%v", urlparams))))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/entries.templ`, Line: 175, Col: 87}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"innerHTML show:top\" hx-target=\"#main-content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 = []any{fmt.Sprintf("%v h-full rounded-2xl modify-bg-color relative group", bgColor)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var5...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var5).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/entries.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><p class=\"text-gray-600 dark:text-gray-300 absolute text-5xl font-medium top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 transition ease-in-out delay-150 group-hover:scale-150 duration-300\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(content)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/entries.templ`, Line: 178, Col: 14}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></a></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

var bgColors = []string{
	"bg-slate-300", "bg-stone-300", "bg-pink-300", "bg-rose-300",
	"bg-orange-300", "bg-amber-300", "bg-lime-300", "bg-green-300",
	"bg-teal-300", "bg-blue-300", "bg-indigo-300", "bg-violet-300", "bg-purple-300",
}

// baseurlparam - suffix is "param="
func individualEntryCards(baseurlparam string, contents []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var8 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var8 == nil {
			templ_7745c5c3_Var8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"px-6 py-6\"><div class=\"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-x-6 gap-y-4 justify-items-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := range contents {
			templ_7745c5c3_Err = individualEntryCard(bgColors[i%len(bgColors)], contents[i], fmt.Sprintf("%v%v", baseurlparam, contents[i])).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func noHealthEntries() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col h-[60vh] items-center justify-center\"><svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" version=\"1.1\" width=\"256\" height=\"256\" viewBox=\"0 0 256 256\" xml:space=\"preserve\"><defs></defs> <g style=\"stroke: none; stroke-width: 0; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: none; fill-rule: nonzero; opacity: 1;\" transform=\"translate(9.164202334630346 9.164202334630318) scale(2.63 2.63)\"><path d=\"M 89.92 26.517 c -0.104 -0.244 -0.301 -0.437 -0.547 -0.536 l -12.914 -5.197 l 1.874 -9.237 c 0.109 -0.541 -0.24 -1.069 -0.781 -1.179 L 27.447 0.204 c -0.538 -0.111 -1.069 0.24 -1.179 0.781 l -2.064 10.173 H 1 c -0.391 0 -0.746 0.228 -0.909 0.583 c -0.163 0.355 -0.104 0.773 0.151 1.069 c 7.078 8.219 10.125 31.67 8.581 66.035 c -0.012 0.272 0.087 0.539 0.276 0.736 c 0.188 0.197 0.45 0.309 0.723 0.309 h 29.407 l 24.484 9.854 c 0.119 0.048 0.246 0.072 0.373 0.072 c 0.134 0 0.267 -0.026 0.392 -0.08 c 0.244 -0.104 0.438 -0.301 0.536 -0.547 l 24.914 -61.907 C 90.026 27.036 90.024 26.761 89.92 26.517 z M 28.029 2.363 l 48.145 9.767 l -1.84 9.07 c 0 0.001 -0.001 0.001 -0.001 0.002 L 63.305 75.57 l -1.814 -0.369 c 0.779 -21.785 0 -49.007 -7.429 -61.202 c -0.012 -0.02 -0.023 -0.04 -0.035 -0.06 c -0.068 -0.111 -0.135 -0.224 -0.204 -0.332 c 0 0 0 -0.001 0 -0.001 c -0.15 -0.236 -0.305 -0.457 -0.46 -0.678 l 0 0 c -0.023 -0.033 -0.045 -0.068 -0.068 -0.101 c -0.044 -0.063 -0.087 -0.132 -0.132 -0.194 c -0.029 -0.04 -0.057 -0.081 -0.086 -0.121 c -0.001 -0.002 -0.001 -0.003 -0.003 -0.005 c -0.259 -0.352 -0.524 -0.687 -0.796 -1.002 c -0.095 -0.11 -0.209 -0.191 -0.332 -0.249 c -0.029 -0.014 -0.062 -0.014 -0.092 -0.025 c -0.099 -0.035 -0.198 -0.063 -0.302 -0.066 c -0.011 0 -0.021 -0.007 -0.032 -0.007 H 26.245 L 28.029 2.363 z M 3.009 13.158 h 22.012 h 26.015 c 0.087 0.108 0.175 0.216 0.261 0.329 c 0.145 0.19 0.29 0.382 0.431 0.586 l -0.005 0.004 c 0.083 0.119 0.165 0.243 0.248 0.364 c 0.138 0.21 0.275 0.423 0.41 0.646 c 0.005 0.009 0.006 0.018 0.011 0.027 l 0.085 0.141 c 0.102 0.168 0.203 0.336 0.301 0.514 c 0.089 0.16 0.174 0.329 0.261 0.495 c 0.003 0.005 0.003 0.011 0.006 0.017 l 0.069 0.132 c 0.095 0.179 0.188 0.358 0.28 0.546 c 0.001 0.002 0.003 0.003 0.004 0.005 c 0.088 0.181 0.173 0.372 0.259 0.56 c 0.001 0.003 0.001 0.005 0.002 0.008 l 0.078 0.169 c 0.078 0.168 0.155 0.336 0.23 0.51 c 0.104 0.241 0.206 0.49 0.308 0.741 l -0.002 0.001 c 0.08 0.196 0.16 0.394 0.238 0.596 c 0.105 0.273 0.209 0.549 0.311 0.833 l -0.004 0.002 c 0.07 0.192 0.138 0.386 0.206 0.582 c 0.1 0.293 0.201 0.584 0.298 0.888 l -0.006 0.002 c 0.075 0.231 0.148 0.463 0.22 0.7 c 0.084 0.275 0.169 0.545 0.25 0.829 c 0.003 0.011 -0.001 0.021 0.002 0.031 l 0.208 0.739 c 0.081 0.298 0.162 0.594 0.241 0.901 c 0.003 0.012 0.001 0.024 0.004 0.036 l 0.225 0.901 c 0.001 0.003 0.003 0.005 0.004 0.008 c 0.065 0.269 0.129 0.539 0.192 0.814 c 0.002 0.012 0.001 0.023 0.004 0.035 l 0.207 0.931 c 0.001 0.006 0.005 0.011 0.007 0.017 c 0.06 0.28 0.119 0.564 0.177 0.85 c 0.002 0.013 0.001 0.025 0.004 0.038 l 0.189 0.958 c 0.002 0.01 0.009 0.018 0.011 0.028 c 0.055 0.286 0.107 0.579 0.16 0.87 c 0.002 0.016 0.001 0.032 0.004 0.048 c 0.06 0.325 0.117 0.653 0.173 0.983 c 0.002 0.014 0.011 0.025 0.014 0.039 c 0.05 0.294 0.097 0.597 0.145 0.896 c 0.002 0.017 0.001 0.033 0.004 0.05 l 0.157 1.015 c 0.003 0.017 0.012 0.029 0.016 0.046 c 0.044 0.293 0.085 0.597 0.127 0.894 c 0.002 0.027 0.002 0.054 0.006 0.081 c 0.05 0.344 0.097 0.69 0.144 1.039 c 0.002 0.016 0.011 0.028 0.014 0.044 c 0.041 0.305 0.078 0.62 0.117 0.93 c 0.001 0.029 0.002 0.057 0.005 0.086 l 0.127 1.06 c 0.002 0.017 0.012 0.03 0.014 0.046 c 0.036 0.308 0.069 0.625 0.104 0.938 c 0 0 0 0 0 0 c 0 0.036 0.002 0.072 0.006 0.108 l 0.113 1.085 c 0.002 0.017 0.011 0.03 0.014 0.047 c 0.033 0.323 0.063 0.652 0.094 0.979 c 0 0.004 -0.002 0.007 -0.002 0.01 c 0 0.032 0.002 0.064 0.005 0.097 l 0.099 1.087 c 0.001 0.017 0.011 0.03 0.013 0.047 c 0.031 0.359 0.061 0.725 0.091 1.09 c 0.001 0.012 -0.005 0.023 -0.004 0.036 l 0.085 1.108 c 0.001 0.018 0.011 0.032 0.013 0.05 c 0.026 0.352 0.051 0.71 0.075 1.067 c 0 0.007 -0.004 0.012 -0.004 0.019 c 0 0.022 0.001 0.046 0.002 0.069 l 0.072 1.122 c 0.001 0.018 0.011 0.032 0.013 0.05 c 0.023 0.368 0.044 0.741 0.065 1.113 c 0 0.007 -0.004 0.013 -0.004 0.02 c 0 0.019 0 0.037 0.002 0.057 l 0.058 1.113 c 0.001 0.017 0.01 0.032 0.012 0.049 c 0.02 0.389 0.038 0.783 0.056 1.177 c 0 0.008 -0.004 0.014 -0.004 0.022 c 0 0.015 0 0.03 0.001 0.045 l 0.046 1.104 c 0.001 0.017 0.01 0.03 0.011 0.047 c 0.016 0.407 0.031 0.819 0.045 1.231 c 0 0.008 -0.004 0.014 -0.004 0.022 c 0 0.012 0 0.023 0.001 0.035 l 0.033 1.101 c 0.001 0.017 0.01 0.031 0.011 0.048 c 0.013 0.425 0.023 0.856 0.034 1.286 c 0 0.007 -0.004 0.013 -0.004 0.02 c 0 0.009 0 0.017 0 0.025 l 0.023 1.074 c 0 0.015 0.008 0.027 0.009 0.042 c 0.009 0.451 0.015 0.91 0.023 1.367 c 0 0.006 -0.004 0.012 -0.004 0.018 c 0 0.005 0 0.01 0 0.015 l 0.013 1.003 c 0 0 0 0.021 0 0.022 c 0 0.015 0.008 0.027 0.009 0.042 c 0.006 0.479 0.007 0.969 0.011 1.454 c 0 0.005 -0.003 0.009 -0.003 0.014 c 0 0.004 0 0.007 0 0.01 l 0.004 0.979 c 0 0.011 0.006 0.02 0.006 0.031 c 0.002 0.51 0 1.03 -0.001 1.547 c 0 0.003 -0.002 0.005 -0.002 0.008 l -0.004 0.883 c 0 0.001 0 0.002 0 0.003 c 0 0.007 0.004 0.013 0.004 0.02 c -0.003 0.559 -0.008 1.125 -0.013 1.691 c 0 0.001 -0.001 0.002 -0.001 0.003 l -0.008 0.605 c 0 0.003 0 0.007 0 0.01 c 0 0.003 0.002 0.006 0.002 0.01 c -0.014 1.229 -0.033 2.47 -0.061 3.734 c -0.003 0.114 -0.005 0.228 -0.008 0.342 c -0.018 0.794 -0.04 1.595 -0.063 2.402 c -0.009 0.302 -0.019 0.607 -0.029 0.911 c -0.02 0.633 -0.042 1.27 -0.065 1.912 c -0.03 0.801 -0.059 1.603 -0.094 2.416 h -19.96 H 10.867 C 11.688 58.434 11.599 25.962 3.009 13.158 z M 63.531 87.516 l -18.945 -7.625 h 15.754 c 0.535 0 0.976 -0.421 0.999 -0.955 c 0.025 -0.57 0.05 -1.139 0.073 -1.708 l 2.474 0.502 c 0.066 0.013 0.133 0.02 0.199 0.02 c 0.044 0 0.086 -0.013 0.13 -0.019 c 0.034 -0.004 0.066 -0.007 0.099 -0.014 c 0.114 -0.027 0.225 -0.067 0.324 -0.133 c 0.221 -0.146 0.375 -0.375 0.428 -0.635 l 10.989 -54.171 l 11.644 4.686 L 63.531 87.516 z\" style=\"stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(96,165,250); fill-rule: nonzero; opacity: 1;\" transform=\" matrix(1 0 0 1 0 0) \" stroke-linecap=\"round\"></path></g></svg> <span class=\"text-lg font-bold text-center text-gray-700 dark:text-gray-100\">No entries found</span></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func HealthEntryDashboard(contents []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var11 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"px-2 flex flex-col gap-5\"><div class=\"flex items-center font-semibold space-x-2\"><a class=\"text-2xl md:text-4xl text-blue-400\" hx-get=\"/entry\" hx-swap=\"innerHTML show:top\" hx-target=\"#main-content\">Years</a><p class=\"text-lg sm:text-xl\">></p></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(contents) == 0 {
				templ_7745c5c3_Err = noHealthEntries().Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = individualEntryCards("year=", contents).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = base("").Render(templ.WithChildren(ctx, templ_7745c5c3_Var11), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func modifyBGColors() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_modifyBGColors_b553`,
		Function: `function __templ_modifyBGColors_b553(){if (localStorage.fitTheme === "light") {
        themeToggle.innerHTML = lightModeHTML
        modifyColorsBasedOnTheme(false)
    } else {
        themeToggle.innerHTML = darkModeHTML
        modifyColorsBasedOnTheme(true)
    }
}`,
		Call:       templ.SafeScript(`__templ_modifyBGColors_b553`),
		CallInline: templ.SafeScriptInline(`__templ_modifyBGColors_b553`),
	}
}

func HealthEntryYears(contents []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var12 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var12 == nil {
			templ_7745c5c3_Var12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title>Fit</title><main class=\"px-2 flex flex-col gap-5\"><div class=\"flex items-center font-semibold space-x-2\"><a class=\"text-2xl md:text-4xl text-blue-400\" hx-get=\"/entry\" hx-swap=\"innerHTML show:top\" hx-target=\"#main-content\">Years</a><p class=\"text-lg sm:text-xl\">></p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(contents) == 0 {
			templ_7745c5c3_Err = noHealthEntries().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = individualEntryCards("year=", contents).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(contents) != 0 {
			templ_7745c5c3_Err = modifyBGColors().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func HealthEntryMonths(year string, contents []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var13 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var13 == nil {
			templ_7745c5c3_Var13 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title>Fit</title><main class=\"px-2 flex flex-col gap-5\"><div class=\"flex items-center font-semibold space-x-2\"><a class=\"text-2xl md:text-4xl text-blue-400\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(string(templ.URL(fmt.Sprintf("/entry?year=%v", year))))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/entries.templ`, Line: 266, Col: 112}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"innerHTML show:top\" hx-target=\"#main-content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(year)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/entries.templ`, Line: 267, Col: 10}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a><p class=\"text-lg sm:text-xl\">></p><p class=\"text-2xl md:text-4xl text-blue-400\">Months</p><p class=\"text-lg sm:text-xl\">></p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(contents) == 0 {
			templ_7745c5c3_Err = noHealthEntries().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = individualEntryCards(fmt.Sprintf("year=%v&month=", year), contents).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(contents) != 0 {
			templ_7745c5c3_Err = modifyBGColors().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func HealthEntryDays(year, month string, contents []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var16 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var16 == nil {
			templ_7745c5c3_Var16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title>Fit</title><main class=\"px-2 flex flex-col gap-5\"><div class=\"flex items-center font-semibold space-x-2\"><p class=\"text-2xl md:text-4xl text-blue-400\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(string(templ.URL(fmt.Sprintf("/entry?year=%v", year))))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/entries.templ`, Line: 290, Col: 112}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"innerHTML show:top\" hx-target=\"#main-content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var18 string
		templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(year)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/entries.templ`, Line: 291, Col: 10}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"text-lg sm:text-xl\">></p><p class=\"text-2xl md:text-4xl text-blue-400\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var19 string
		templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(string(templ.URL(fmt.Sprintf("/entry?year=%v&month=%v", year, month))))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/entries.templ`, Line: 294, Col: 128}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"innerHTML show:top\" hx-target=\"#main-content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var20 string
		templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(month)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/entries.templ`, Line: 295, Col: 11}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"text-lg sm:text-xl\">></p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(contents) == 0 {
			templ_7745c5c3_Err = noHealthEntries().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = individualEntryCards(fmt.Sprintf("year=%v&month=%v&day=", year, month), contents).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(contents) != 0 {
			templ_7745c5c3_Err = modifyBGColors().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
