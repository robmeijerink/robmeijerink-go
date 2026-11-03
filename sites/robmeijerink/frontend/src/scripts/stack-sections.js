export function initStackSections() {
    const sheets = document.querySelectorAll('[data-sheet]')
    const query = window.matchMedia('(min-width: 64rem) and (prefers-reduced-motion: no-preference)')

    if (!sheets.length) return

    function apply() {
        sheets.forEach((sheet) => {
            if (!query.matches) {
                sheet.classList.remove('is-sheet')
                sheet.style.removeProperty('top')
                return
            }

            const top = Math.min(0, window.innerHeight - sheet.offsetHeight)

            sheet.style.top = `${top}px`
            sheet.classList.add('is-sheet')
        })
    }

    apply()

    window.addEventListener('resize', apply)
    query.addEventListener('change', apply)

    if ('ResizeObserver' in window) {
        const observer = new ResizeObserver(apply)
        sheets.forEach((sheet) => observer.observe(sheet))
    }
}
