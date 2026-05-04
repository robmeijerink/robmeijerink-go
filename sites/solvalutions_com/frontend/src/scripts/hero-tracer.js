export function initHeroTracer(canvas) {
    if (!canvas) return

    const ctx = canvas.getContext('2d')

    let width, height, columns, rows
    const gridSize = 40
    const tracers = []
    const numTracers = 8

    function resize() {
        width = canvas.width = window.innerWidth
        height = canvas.height = window.innerHeight
        columns = Math.floor(width / gridSize)
        rows = Math.floor(height / gridSize)
    }

    window.addEventListener('resize', resize)
    resize()

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

            if (this.x % gridSize === 0 && this.y % gridSize === 0) {
                if (Math.random() < 0.3) {
                    const turn = Math.random() < 0.5 ? -1 : 1
                    this.dir = (this.dir + turn + 4) % 4
                }
            }

            if (this.x < 0 || this.x > width || this.y < 0 || this.y > height) {
                this.init()
            }
        }

        draw() {
            if (this.history.length < 2) return

            ctx.beginPath()
            ctx.lineWidth = 2
            ctx.lineCap = 'round'
            ctx.strokeStyle = '#38bdf8'

            for (let i = 0; i < this.history.length - 1; i++) {
                const opacity = i / this.history.length
                ctx.globalAlpha = opacity
                ctx.moveTo(this.history[i].x, this.history[i].y)
                ctx.lineTo(this.history[i+1].x, this.history[i+1].y)
            }
            ctx.stroke()

            const head = this.history[this.history.length - 1]
            ctx.globalAlpha = 1
            ctx.fillStyle = '#fff'
            ctx.shadowBlur = 10
            ctx.shadowColor = '#38bdf8'
            ctx.fillRect(head.x - 2, head.y - 2, 4, 4)
            ctx.shadowBlur = 0
        }
    }

    for (let i = 0; i < numTracers; i++) {
        tracers.push(new Tracer())
    }

    function animate() {
        ctx.clearRect(0, 0, width, height)

        tracers.forEach(t => {
            t.update()
            t.draw()
        })

        requestAnimationFrame(animate)
    }

    animate()
}

