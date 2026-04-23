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

const openBtn = document.getElementById('mobile-menu-open')
const closeBtn = document.getElementById('mobile-menu-close')
const menu = document.getElementById('mobile-menu')
const backdrop = document.getElementById('mobile-menu-backdrop')

function toggleMenu(isOpen) {
    if (isOpen) {
        menu.classList.remove('translate-x-full')
        backdrop.classList.add('opacity-100', 'pointer-events-auto')
        backdrop.classList.remove('opacity-0', 'pointer-events-none')
        document.body.style.overflow = 'hidden'
    } else {
        menu.classList.add('translate-x-full')
        backdrop.classList.remove('opacity-100', 'pointer-events-auto')
        backdrop.classList.add('opacity-0', 'pointer-events-none')
        document.body.style.overflow = ''
    }
}

openBtn.addEventListener('click', () => toggleMenu(true))
closeBtn.addEventListener('click', () => toggleMenu(false))
backdrop.addEventListener('click', () => toggleMenu(false))
