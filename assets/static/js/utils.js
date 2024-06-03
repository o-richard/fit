const darkModeHTML = `
    <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" width="40" height="40" viewBox="0 0 256 256" xml:space="preserve">
        <defs> </defs>
        <g style="stroke: none; stroke-width: 0; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: none; fill-rule: nonzero; opacity: 1;" transform="translate(45.02412451361867 45.024124513618645) scale(1.83 1.83)" >
            <path d="M 46.715 90 c -3.908 0 -7.841 -0.514 -11.717 -1.552 C 23.391 85.337 13.69 77.893 7.682 67.487 C 1.674 57.08 0.077 44.957 3.188 33.349 c 3.11 -11.607 10.554 -21.308 20.961 -27.316 c 8.601 -4.967 18.349 -6.923 28.193 -5.659 c 1.257 0.162 2.277 1.095 2.548 2.332 c 0.271 1.238 -0.265 2.512 -1.338 3.185 c -13.943 8.735 -18.418 26.742 -10.188 40.996 l 0 0 C 51.592 61.14 69.426 66.268 83.96 58.56 c 1.117 -0.596 2.491 -0.421 3.426 0.434 c 0.936 0.854 1.235 2.204 0.746 3.373 c -3.826 9.156 -10.395 16.621 -18.997 21.586 C 62.204 87.955 54.509 90 46.715 90 z M 43.74 6.101 c -5.805 0.421 -11.436 2.15 -16.592 5.127 c -9.019 5.207 -15.47 13.614 -18.166 23.674 C 6.287 44.961 7.67 55.469 12.877 64.488 c 5.207 9.019 13.614 15.471 23.673 18.165 c 10.058 2.697 20.567 1.311 29.585 -3.895 c 5.156 -2.977 9.47 -6.989 12.737 -11.806 c -15.547 4.094 -32.303 -2.515 -40.705 -17.066 l 0 0 C 29.768 35.336 32.427 17.518 43.74 6.101 z" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(59,130,246); fill-rule: nonzero; opacity: 1;" transform=" matrix(1 0 0 1 0 0) " stroke-linecap="round" />
        </g>
    </svg>
`
const lightModeHTML = `
    <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" width="40" height="40" viewBox="0 0 256 256" xml:space="preserve">
        <defs> </defs>
        <g style="stroke: none; stroke-width: 0; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: none; fill-rule: nonzero; opacity: 1;" transform="translate(18.12918287937744 18.12918287937741) scale(2.43 2.43)" >
            <path d="M 45 68 c -12.682 0 -23 -10.317 -23 -23 c 0 -12.682 10.318 -23 23 -23 c 12.683 0 23 10.318 23 23 C 68 57.683 57.683 68 45 68 z M 45 28 c -9.374 0 -17 7.626 -17 17 s 7.626 17 17 17 s 17 -7.626 17 -17 S 54.374 28 45 28 z" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(59,130,246); fill-rule: nonzero; opacity: 1;" transform=" matrix(1 0 0 1 0 0) " stroke-linecap="round" />
            <path d="M 45 17.556 c -1.657 0 -3 -1.343 -3 -3 V 3 c 0 -1.657 1.343 -3 3 -3 c 1.657 0 3 1.343 3 3 v 11.556 C 48 16.212 46.657 17.556 45 17.556 z" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(59,130,246); fill-rule: nonzero; opacity: 1;" transform=" matrix(1 0 0 1 0 0) " stroke-linecap="round" />
            <path d="M 45 90 c -1.657 0 -3 -1.343 -3 -3 V 75.444 c 0 -1.657 1.343 -3 3 -3 c 1.657 0 3 1.343 3 3 V 87 C 48 88.657 46.657 90 45 90 z" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(59,130,246); fill-rule: nonzero; opacity: 1;" transform=" matrix(1 0 0 1 0 0) " stroke-linecap="round" />
            <path d="M 14.556 48 H 3 c -1.657 0 -3 -1.343 -3 -3 c 0 -1.657 1.343 -3 3 -3 h 11.556 c 1.657 0 3 1.343 3 3 C 17.556 46.657 16.212 48 14.556 48 z" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(59,130,246); fill-rule: nonzero; opacity: 1;" transform=" matrix(1 0 0 1 0 0) " stroke-linecap="round" />
            <path d="M 87 48 H 75.444 c -1.657 0 -3 -1.343 -3 -3 c 0 -1.657 1.343 -3 3 -3 H 87 c 1.657 0 3 1.343 3 3 C 90 46.657 88.657 48 87 48 z" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(59,130,246); fill-rule: nonzero; opacity: 1;" transform=" matrix(1 0 0 1 0 0) " stroke-linecap="round" />
            <path d="M 66.527 26.473 c -0.768 0 -1.535 -0.293 -2.121 -0.878 c -1.172 -1.172 -1.172 -3.071 0 -4.243 l 8.171 -8.171 c 1.172 -1.172 3.07 -1.171 4.242 0 c 1.172 1.172 1.172 3.071 0 4.243 l -8.171 8.171 C 68.063 26.18 67.295 26.473 66.527 26.473 z" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(59,130,246); fill-rule: nonzero; opacity: 1;" transform=" matrix(1 0 0 1 0 0) " stroke-linecap="round" />
            <path d="M 15.302 77.698 c -0.768 0 -1.536 -0.293 -2.121 -0.879 c -1.172 -1.171 -1.172 -3.071 0 -4.242 l 8.171 -8.171 c 1.171 -1.172 3.071 -1.172 4.242 0 c 1.172 1.171 1.172 3.071 0 4.242 l -8.171 8.171 C 16.837 77.405 16.069 77.698 15.302 77.698 z" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(59,130,246); fill-rule: nonzero; opacity: 1;" transform=" matrix(1 0 0 1 0 0) " stroke-linecap="round" />
            <path d="M 23.473 26.473 c -0.768 0 -1.536 -0.293 -2.121 -0.878 l -8.171 -8.171 c -1.172 -1.172 -1.172 -3.071 0 -4.243 c 1.172 -1.172 3.072 -1.171 4.243 0 l 8.171 8.171 c 1.172 1.172 1.172 3.071 0 4.243 C 25.008 26.18 24.24 26.473 23.473 26.473 z" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(59,130,246); fill-rule: nonzero; opacity: 1;" transform=" matrix(1 0 0 1 0 0) " stroke-linecap="round" />
            <path d="M 74.698 77.698 c -0.768 0 -1.535 -0.293 -2.121 -0.879 l -8.171 -8.171 c -1.172 -1.171 -1.172 -3.071 0 -4.242 c 1.172 -1.172 3.07 -1.172 4.242 0 l 8.171 8.171 c 1.172 1.171 1.172 3.071 0 4.242 C 76.233 77.405 75.466 77.698 74.698 77.698 z" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(59,130,246); fill-rule: nonzero; opacity: 1;" transform=" matrix(1 0 0 1 0 0) " stroke-linecap="round" />
        </g>
    </svg>
`
const themeToggle = document.getElementById("themeToggle")
if (localStorage.fitTheme === "light") {
    themeToggle.innerHTML = lightModeHTML
    modifyColorsBasedOnTheme(false)
} else {
    themeToggle.innerHTML = darkModeHTML
    modifyColorsBasedOnTheme(true)
}

function changeTheme() {
    document.documentElement.classList.toggle('dark')
    localStorage.fitTheme = (document.documentElement.classList.contains('dark')) ? 'dark' : 'light'
    if (localStorage.fitTheme === "light") {
        themeToggle.innerHTML = lightModeHTML
        modifyColorsBasedOnTheme(false)
    } else {
        themeToggle.innerHTML = darkModeHTML
        modifyColorsBasedOnTheme(true)
    }
}
function modifyColorsBasedOnTheme(isDarkMode) {
    const elements = document.querySelectorAll(".modify-bg-color")
    const darkModeReplacements = {
        "bg-slate-300": "bg-slate-600", "bg-stone-300": "bg-stone-600", "bg-pink-300": "bg-pink-600", "bg-rose-300": "bg-rose-600",
        "bg-orange-300": "bg-orange-600", "bg-amber-300": "bg-amber-600", "bg-lime-300": "bg-lime-600", "bg-green-300": "bg-green-600", "bg-teal-300": "bg-teal-600",
        "bg-blue-300": "bg-blue-600", "bg-indigo-300": "bg-indigo-600", "bg-violet-300": "bg-violet-600", "bg-purple-300": "bg-purple-600",
    }
    const lightModeReplacements = {
        "bg-slate-600": "bg-slate-300", "bg-stone-600": "bg-stone-300", "bg-pink-600": "bg-pink-300", "bg-rose-600": "bg-rose-300",
        "bg-orange-600": "bg-orange-300", "bg-amber-600": "bg-amber-300", "bg-lime-600": "bg-lime-300", "bg-green-600": "bg-green-300", "bg-teal-600": "bg-teal-300",
        "bg-blue-600": "bg-blue-300", "bg-indigo-600": "bg-indigo-300", "bg-violet-600": "bg-violet-300", "bg-purple-600": "bg-purple-300",
    }

    for (let i = 0; i < elements.length; i++) {
        const classes = elements[i].classList
        if (isDarkMode) {
            for (const [key, value] of Object.entries(darkModeReplacements)) {
                if (classes.contains(key)) {
                    elements[i].classList.replace(key, value)
                    break
                }
            }
        } else {
            for (const [key, value] of Object.entries(lightModeReplacements)) {
                if (classes.contains(key)) {
                    elements[i].classList.replace(key, value)
                    break
                }
            }
        }
    }
}

toastr.options = {
    "closeButton": false,
    "debug": false,
    "newestOnTop": false,
    "progressBar": true,
    "positionClass": "toast-top-right",
    "preventDuplicates": false,
    "onclick": null,
    "showDuration": "300",
    "hideDuration": "1000",
    "timeOut": "5000",
    "extendedTimeOut": "1000",
    "showEasing": "swing",
    "hideEasing": "linear",
    "showMethod": "fadeIn",
    "hideMethod": "fadeOut"
}