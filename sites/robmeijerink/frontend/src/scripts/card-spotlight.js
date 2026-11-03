export function initCardSpotlight() {
    const cards = document.querySelectorAll('[data-spotlight]')

    if (!cards.length) return
    if (!window.matchMedia('(hover: hover) and (pointer: fine)').matches) return
    if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) return

    cards.forEach((card) => {
        let frame = null
        let x = 0
        let y = 0

        function update() {
            card.style.setProperty('--mx', `${x}px`)
            card.style.setProperty('--my', `${y}px`)
            frame = null
        }

        card.addEventListener('pointermove', (event) => {
            if (event.pointerType !== 'mouse') return

            const rect = card.getBoundingClientRect()

            x = event.clientX - rect.left
            y = event.clientY - rect.top

            if (frame === null) frame = requestAnimationFrame(update)
        })

        card.addEventListener('pointerenter', (event) => {
            if (event.pointerType === 'mouse') card.style.setProperty('--spot', '1')
        })

        card.addEventListener('pointerleave', () => {
            card.style.setProperty('--spot', '0')
        })
    })
}
