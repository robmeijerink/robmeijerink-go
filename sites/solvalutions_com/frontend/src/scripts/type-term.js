const CHAR_MS = 22
const LINE_GAP_MS = 120

export function initTypeTerm() {
    const terms = document.querySelectorAll('[data-type]')
    if (!terms.length) return

    const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches
    if (reduced || !('IntersectionObserver' in window)) return

    terms.forEach(term => {
        term.querySelectorAll('.tl').forEach(line => {
            line.style.clipPath = 'inset(0 100% 0 0)'
        })
    })

    const io = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (!entry.isIntersecting) return
            io.unobserve(entry.target)
            play(entry.target)
        })
    }, { threshold: 0.35 })

    terms.forEach(term => io.observe(term))
}

function play(term) {
    const lines = Array.from(term.querySelectorAll('.tl'))
    const charMs = Number(term.dataset.speed || CHAR_MS)
    let t = Number(term.dataset.delay || 0)

    lines.forEach(line => {
        const chars = Math.max(line.textContent.length, 1)
        const dur = line.hasAttribute('data-instant') ? 0 : chars * charMs

        setTimeout(() => {
            line.style.transition = dur ? `clip-path ${dur}ms steps(${chars}, end)` : 'none'
            line.style.clipPath = 'inset(0 0 0 0)'
        }, t)

        t += dur + LINE_GAP_MS
    })

    setTimeout(() => {
        lines.forEach(line => { line.style.clipPath = 'none' })
    }, t + 100)
}
