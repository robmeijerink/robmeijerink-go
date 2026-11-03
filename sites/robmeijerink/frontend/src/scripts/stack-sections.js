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

    function targetTop(target) {
        const sheet = target.closest('[data-sheet]')
        const inside = target === sheet ? 0 : target.getBoundingClientRect().top - sheet.getBoundingClientRect().top
        const offset = parseFloat(getComputedStyle(target).scrollMarginTop) || 0

        const wanted = flowTop(sheet) + inside - offset
        const pinned = flowTop(sheet) + sheet.offsetHeight - window.innerHeight

        return sheet.offsetHeight > window.innerHeight ? Math.min(wanted, pinned) : wanted
    }

    function markTarget(target) {
        document.querySelectorAll('.is-target').forEach((node) => node.classList.remove('is-target'))
        target.classList.add('is-target')
    }

    function scrollToHash() {
        if (!query.matches || !window.location.hash) return

        const target = document.getElementById(window.location.hash.slice(1))

        if (!target || !target.closest('[data-sheet]')) return

        window.scrollTo({ top: targetTop(target), behavior: 'auto' })
    }

    document.addEventListener('click', (event) => {
        if (!query.matches || event.defaultPrevented || event.button !== 0) return
        if (event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return

        const link = event.target.closest('a[href^="#"]')

        if (!link) return

        const id = link.getAttribute('href').slice(1)
        const target = id ? document.getElementById(id) : null

        if (!target || !target.closest('[data-sheet]')) return

        event.preventDefault()
        markTarget(target)
        window.scrollTo({ top: targetTop(target), behavior: 'smooth' })
        history.pushState(null, '', `#${id}`)
    })

    let interacted = false

    ;['wheel', 'touchstart', 'keydown', 'pointerdown'].forEach((name) => {
        window.addEventListener(name, () => (interacted = true), { passive: true, once: true })
    })

    function settleHash() {
        if (!interacted) scrollToHash()
    }

    requestAnimationFrame(settleHash)
    window.addEventListener('load', settleHash)
    window.addEventListener('pagereveal', settleHash)
    setTimeout(settleHash, 150)
    setTimeout(settleHash, 450)

    if (document.fonts && document.fonts.ready) document.fonts.ready.then(settleHash)
}
