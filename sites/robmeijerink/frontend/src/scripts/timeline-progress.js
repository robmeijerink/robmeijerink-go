export function initTimelineProgress() {
    const timeline = document.querySelector('[data-timeline]')

    if (!timeline) return

    const label = document.querySelector('[data-period-label]')
    const items = Array.from(timeline.querySelectorAll('[data-reveal]'))
    const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches
    let shown = null

    function showPeriod(viewport) {
        let current = items[0]

        items.forEach((item) => {
            if (item.getBoundingClientRect().top <= viewport * 0.55) current = item
        })

        const text = current.querySelector('span').textContent.trim()

        if (text === shown) return

        shown = text
        label.textContent = text
        label.classList.add('is-live')

        if (!reduced) {
            label.animate([{ opacity: 0, transform: 'translateY(0.5rem)' }, { opacity: 1, transform: 'none' }], { duration: 320, easing: 'ease-out' })
        }
    }

    let frame = null

    function update() {
        const rect = timeline.getBoundingClientRect()
        const viewport = window.innerHeight
        const progress = (viewport * 0.8 - rect.top) / rect.height
        const clamped = Math.min(1, Math.max(0, progress))

        timeline.style.setProperty('--timeline-progress', clamped.toFixed(3))

        if (label) showPeriod(viewport)
        frame = null
    }

    function schedule() {
        if (frame === null) frame = requestAnimationFrame(update)
    }

    window.addEventListener('scroll', schedule, { passive: true })
    window.addEventListener('resize', schedule)
    update()
}
