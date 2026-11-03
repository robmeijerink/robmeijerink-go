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

    function flowTop(sheet) {
        const container = sheet.parentElement
        let top = container.getBoundingClientRect().top + window.scrollY

        for (let node = container.firstElementChild; node && node !== sheet; node = node.nextElementSibling) {
            top += node.offsetHeight
        }

        return top
    }

    document.addEventListener('click', (event) => {
        if (!query.matches || event.defaultPrevented || event.button !== 0) return
        if (event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return

        const link = event.target.closest('a[href^="#"]')

        if (!link) return

        const id = link.getAttribute('href').slice(1)
        const target = id ? document.getElementById(id) : null

        if (!target || !target.hasAttribute('data-sheet')) return

        event.preventDefault()

        const offset = parseFloat(getComputedStyle(target).scrollMarginTop) || 0

        window.scrollTo({ top: flowTop(target) - offset, behavior: 'smooth' })
        history.pushState(null, '', `#${id}`)
    })
}
