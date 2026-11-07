const ICE = '159,211,255'
const MOUSE_RADIUS = 120

export function initConstellation(canvas) {
    if (!canvas) return

    const host = canvas.parentElement
    const ctx = canvas.getContext('2d')
    const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches
    const mouse = { x: -9999, y: -9999 }

    let width = 0
    let height = 0
    let connectDist = 110
    let particles = []
    let running = false
    let raf = 0

    function build() {
        const count = Math.min(Math.max(Math.round((width * height) / 7000), 20), 80)
        connectDist = width < 480 ? 100 : 140
        particles = []

        for (let i = 0; i < count; i++) {
            particles.push({
                x: Math.random() * width,
                y: Math.random() * height,
                vx: (Math.random() - 0.5) * 0.5,
                vy: (Math.random() - 0.5) * 0.5,
                r: Math.random() * 1.8 + 1.4,
                glow: 0,
            })
        }
    }

    function resize() {
        const rect = host.getBoundingClientRect()
        const dpr = Math.min(window.devicePixelRatio || 1, 2)
        width = rect.width
        height = rect.height
        canvas.width = width * dpr
        canvas.height = height * dpr
        ctx.setTransform(dpr, 0, 0, dpr, 0, 0)
        build()
        draw()
    }

    function step() {
        for (const p of particles) {
            p.x += p.vx
            p.y += p.vy

            if (p.x < -10) p.x = width + 10
            if (p.x > width + 10) p.x = -10
            if (p.y < -10) p.y = height + 10
            if (p.y > height + 10) p.y = -10

            const dx = p.x - mouse.x
            const dy = p.y - mouse.y
            const d = Math.hypot(dx, dy)
            let target = 0

            if (d < MOUSE_RADIUS) {
                target = 1 - d / MOUSE_RADIUS
                p.x += (dx / (d || 1)) * target * 1.2
                p.y += (dy / (d || 1)) * target * 1.2
            }

            p.glow += (target - p.glow) * 0.15
        }
    }

    function draw() {
        ctx.clearRect(0, 0, width, height)

        if (mouse.x > 0) {
            const grad = ctx.createRadialGradient(mouse.x, mouse.y, 0, mouse.x, mouse.y, MOUSE_RADIUS)
            grad.addColorStop(0, `rgba(${ICE},0.08)`)
            grad.addColorStop(1, `rgba(${ICE},0)`)
            ctx.fillStyle = grad
            ctx.fillRect(mouse.x - MOUSE_RADIUS, mouse.y - MOUSE_RADIUS, MOUSE_RADIUS * 2, MOUSE_RADIUS * 2)
        }

        for (let i = 0; i < particles.length; i++) {
            const a = particles[i]

            for (let j = i + 1; j < particles.length; j++) {
                const b = particles[j]
                const d = Math.hypot(a.x - b.x, a.y - b.y)
                if (d >= connectDist) continue

                const glow = Math.max(a.glow, b.glow)
                ctx.strokeStyle = `rgba(${ICE},${(1 - d / connectDist) * 0.6 + glow * 0.3})`
                ctx.lineWidth = 1 + glow * 0.6
                ctx.beginPath()
                ctx.moveTo(a.x, a.y)
                ctx.lineTo(b.x, b.y)
                ctx.stroke()
            }
        }

        for (const p of particles) {
            ctx.fillStyle = `rgba(${ICE},${0.12 + p.glow * 0.2})`
            ctx.beginPath()
            ctx.arc(p.x, p.y, p.r * 3 + p.glow * 4, 0, Math.PI * 2)
            ctx.fill()

            ctx.fillStyle = `rgba(${ICE},${0.85 + p.glow * 0.15})`
            ctx.beginPath()
            ctx.arc(p.x, p.y, p.r + p.glow * 2.2, 0, Math.PI * 2)
            ctx.fill()
        }
    }

    function loop() {
        if (!running) return
        step()
        draw()
        raf = requestAnimationFrame(loop)
    }

    function start() {
        if (running || reduced) return
        running = true
        raf = requestAnimationFrame(loop)
    }

    function stop() {
        running = false
        cancelAnimationFrame(raf)
    }

    const zone = canvas.closest('[data-constellation-zone]') || host

    zone.addEventListener('pointermove', (e) => {
        const rect = host.getBoundingClientRect()
        mouse.x = e.clientX - rect.left
        mouse.y = e.clientY - rect.top
    })

    zone.addEventListener('pointerleave', () => {
        mouse.x = -9999
        mouse.y = -9999
    })

    new ResizeObserver(resize).observe(host)

    new IntersectionObserver(([entry]) => {
        if (entry.isIntersecting && !document.hidden) start()
        else stop()
    }).observe(host)

    document.addEventListener('visibilitychange', () => {
        if (document.hidden) stop()
        else if (!reduced) start()
    })

    resize()
}
