const STEP_MS = 1300

export function initSteps() {
    const root = document.querySelector('[data-steps]')
    if (!root) return

    const steps = Array.from(root.querySelectorAll('[data-step]'))
    const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches
    let timer = 0

    function activate(index) {
        steps.forEach((step, i) => {
            step.dataset.state = i < index ? 'done' : i === index ? 'active' : 'idle'
        })
    }

    steps.forEach((step, i) => {
        for (const event of ['mouseenter', 'focusin', 'click']) {
            step.addEventListener(event, () => {
                clearTimeout(timer)
                activate(i)
            })
        }
    })

    if (reduced) {
        activate(steps.length - 1)
        return
    }

    activate(0)

    if (!('IntersectionObserver' in window)) return

    const io = new IntersectionObserver(([entry]) => {
        if (!entry.isIntersecting) return
        io.disconnect()

        let i = 0
        const next = () => {
            i += 1
            if (i >= steps.length) return
            activate(i)
            timer = setTimeout(next, STEP_MS)
        }

        timer = setTimeout(next, STEP_MS)
    }, { threshold: 0.4 })

    io.observe(root)
}
