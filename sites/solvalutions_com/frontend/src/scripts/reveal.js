export function initReveal() {
    const items = document.querySelectorAll('.reveal, [data-draw]')
    if (!items.length) return

    if (!('IntersectionObserver' in window)) {
        items.forEach(el => el.classList.add('is-in'))
        return
    }

    const io = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (!entry.isIntersecting) return
            entry.target.classList.add('is-in')
            io.unobserve(entry.target)
        })
    }, { threshold: 0.15, rootMargin: '0px 0px -8% 0px' })

    items.forEach(el => io.observe(el))
}
