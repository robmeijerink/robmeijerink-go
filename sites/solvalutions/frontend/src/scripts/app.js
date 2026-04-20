import '../styles/app.css'

// Main
const tracerCanvas = document.getElementById('tech-grid-tracer-canvas')

if (tracerCanvas) {
    const loadHeroTracer = async () => {
        try {
            (await import('./hero-tracer.js')).initHeroTracer(tracerCanvas)
        } catch (err) {
            console.error("Kon de hero tracer niet laden:", err)
        }
    }

    loadHeroTracer()
}
