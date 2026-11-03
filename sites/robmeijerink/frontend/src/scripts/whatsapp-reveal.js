export function initWhatsappReveal() {
    const wrap = document.querySelector('[data-whatsapp]')
    const hero = document.querySelector('main > section')

    if (!wrap || !hero) return

    if (!('IntersectionObserver' in window)) {
        wrap.classList.add('is-shown')
        return
    }

    const observer = new IntersectionObserver(
        ([entry]) => {
            const passed = !entry.isIntersecting && entry.boundingClientRect.top < 0

            wrap.classList.toggle('is-shown', passed)
        },
        { threshold: 0 },
    )

    observer.observe(hero)
}
