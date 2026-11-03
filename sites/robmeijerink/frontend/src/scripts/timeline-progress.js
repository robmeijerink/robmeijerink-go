export function initTimelineProgress() {
    const timeline = document.querySelector('[data-timeline]')

    if (!timeline) return

    let frame = null

    function update() {
        const rect = timeline.getBoundingClientRect()
        const viewport = window.innerHeight
        const progress = (viewport * 0.8 - rect.top) / rect.height
        const clamped = Math.min(1, Math.max(0, progress))

        timeline.style.setProperty('--timeline-progress', clamped.toFixed(3))
        frame = null
    }

    function schedule() {
        if (frame === null) frame = requestAnimationFrame(update)
    }

    window.addEventListener('scroll', schedule, { passive: true })
    window.addEventListener('resize', schedule)
    update()
}
