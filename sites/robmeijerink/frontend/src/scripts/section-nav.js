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

        frame = null
    }

    function schedule() {
        if (frame === null) frame = requestAnimationFrame(update)
    }

    window.addEventListener('scroll', schedule, { passive: true })
    window.addEventListener('resize', schedule)
    update()
}
