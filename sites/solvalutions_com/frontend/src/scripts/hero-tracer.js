const ICE = '159, 211, 255'

export function initHeroTracer(canvas) {
    if (!canvas) return

    const ctx = canvas.getContext('2d')
    const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches
    const gridSize = 48
    const tracers = []

    let width, height, columns, rows
    let running = false
    let raf = 0

    function resize() {
        const rect = canvas.parentElement.getBoundingClientRect()
        const dpr = Math.min(window.devicePixelRatio || 1, 2)
        width = rect.width
        height = rect.height
        canvas.width = width * dpr
        canvas.height = height * dpr
        canvas.style.width = width + 'px'
        canvas.style.height = height + 'px'
        ctx.setTransform(dpr, 0, 0, dpr, 0, 0)
        columns = Math.floor(width / gridSize)
        rows = Math.floor(height / gridSize)
    }

    class Tracer {
        constructor() {
            this.init()
        }

        init() {
            this.x = Math.floor(Math.random() * columns) * gridSize
            this.y = Math.floor(Math.random() * rows) * gridSize
            this.history = []
            this.maxLength = 15 + Math.random() * 20
            this.speed = 2
            this.dir = Math.floor(Math.random() * 4)
        }

        update() {
            this.history.push({ x: this.x, y: this.y })
            if (this.history.length > this.maxLength) this.history.shift()

            if (this.dir === 0) this.x += this.speed
            if (this.dir === 1) this.y += this.speed
            if (this.dir === 2) this.x -= this.speed
            if (this.dir === 3) this.y -= this.speed

            if (this.x % gridSize === 0 && this.y % gridSize === 0 && Math.random() < 0.3) {
                const turn = Math.random() < 0.5 ? -1 : 1
                this.dir = (this.dir + turn + 4) % 4
            }

            if (this.x < 0 || this.x > width || this.y < 0 || this.y > height) {
                this.init()
            }
        }

        draw() {
            if (this.history.length < 2) return

            ctx.lineWidth = 1.5
            ctx.lineCap = 'round'

            for (let i = 0; i < this.history.length - 1; i++) {
                const a = this.history[i]
                const b = this.history[i + 1]
                ctx.strokeStyle = `rgba(${ICE}, ${(i / this.history.length) * 0.8})`
                ctx.beginPath()
                ctx.moveTo(a.x, a.y)
                ctx.lineTo(b.x, b.y)
                ctx.stroke()
            }

            const head = this.history[this.history.length - 1]
            ctx.fillStyle = '#fff'
            ctx.fillRect(head.x - 2, head.y - 2, 4, 4)
        }
    }

    function build() {
        tracers.length = 0
        const count = width < 768 ? 4 : 8
        for (let i = 0; i < count; i++) tracers.push(new Tracer())
    }

    function frame() {
        ctx.clearRect(0, 0, width, height)

        for (const t of tracers) {
            t.update()
            t.draw()
        }
    }

    function loop() {
        if (!running) return
        frame()
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

    resize()
    build()

    new ResizeObserver(() => {
        resize()
        build()
    }).observe(canvas.parentElement)

    new IntersectionObserver(([entry]) => {
        if (entry.isIntersecting && !document.hidden) start()
        else stop()
    }).observe(canvas)

    document.addEventListener('visibilitychange', () => {
        if (document.hidden) stop()
        else start()
    })
}
