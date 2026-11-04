export function initScrollWords() {
    const targets = document.querySelectorAll('[data-scroll-words]')

    if (!targets.length || window.matchMedia('(prefers-reduced-motion: reduce)').matches) return

    const groups = Array.from(targets).map((target) => ({ target, words: splitWords(target) }))

    function splitWords(root) {
        const words = []
        const walker = document.createTreeWalker(root, NodeFilter.SHOW_TEXT)
        const nodes = []

        while (walker.nextNode()) nodes.push(walker.currentNode)

        nodes.forEach((node) => {
            const fragment = document.createDocumentFragment()

            node.textContent.split(/(\s+)/).forEach((part) => {
                if (!part.trim()) {
                    fragment.append(part)
                    return
                }

                const span = document.createElement('span')
                span.textContent = part
                span.style.display = 'inline-block'
                words.push(span)
                fragment.append(span)
            })

            node.replaceWith(fragment)
        })

        return words
    }

    let frame = null

    function update() {
        const viewport = window.innerHeight

        groups.forEach(({ target, words }) => {
            const rect = target.getBoundingClientRect()
            const progress = Math.min(1, Math.max(0, (viewport * 0.9 - rect.top) / (viewport * 0.5)))

            words.forEach((word, index) => {
                const start = index / words.length
                const local = Math.min(1, Math.max(0, (progress - start * 0.6) / 0.4))

                word.style.opacity = (0.15 + local * 0.85).toFixed(2)
                word.style.filter = `blur(${((1 - local) * 8).toFixed(1)}px)`
            })
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
