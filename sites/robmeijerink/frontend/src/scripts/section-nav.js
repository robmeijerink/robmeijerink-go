export function initSectionNav() {
    const links = document.querySelectorAll('[data-section-link]')

    if (!links.length) return

    const items = Array.from(links)
        .map((link) => ({ link, section: document.getElementById(link.dataset.sectionLink) }))
        .filter((item) => item.section)

    let frame = null

    function update() {
        const threshold = window.innerHeight * 0.4
        let current = -1

        items.forEach((item, index) => {
            if (item.section.getBoundingClientRect().top <= threshold) current = index
        })

        items.forEach((item, index) => {
            if (index === current) {
                item.link.setAttribute('aria-current', 'true')
            } else {
                item.link.removeAttribute('aria-current')
            }

            item.link.dataset.passed = index < current ? 'true' : 'false'
        })

        const active = items[current]
        const scroller = active && active.link.closest('.pipeline-scroll')

        if (scroller && scroller.scrollWidth > scroller.clientWidth) {
            const target = active.link.offsetLeft - (scroller.clientWidth - active.link.clientWidth) / 2

            scroller.scrollTo({ left: target, behavior: 'smooth' })
        }

        frame = null
    }

    function schedule() {
        if (frame === null) frame = requestAnimationFrame(update)
    }

    window.addEventListener('scroll', schedule, { passive: true })
    window.addEventListener('resize', schedule)
    update()
}
