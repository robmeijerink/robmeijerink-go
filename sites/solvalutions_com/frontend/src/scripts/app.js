import '../styles/app.css'
import { initReveal } from './reveal.js'
import { initTypeTerm } from './type-term.js'
import { initSteps } from './steps.js'

initReveal()
initTypeTerm()
initSteps()

const tracerCanvas = document.getElementById('tech-grid-tracer-canvas')

if (tracerCanvas) {
    import('./hero-tracer.js')
        .then(m => m.initHeroTracer(tracerCanvas))
        .catch(err => console.error('Hero tracer failed to load:', err))
}

const constellations = document.querySelectorAll('canvas[data-constellation]')

if (constellations.length) {
    import('./constellation.js')
        .then(m => constellations.forEach(canvas => m.initConstellation(canvas)))
        .catch(err => console.error('Constellation failed to load:', err))
}

const openBtn = document.getElementById('mobile-menu-open')
const closeBtn = document.getElementById('mobile-menu-close')
const menu = document.getElementById('mobile-menu')
const backdrop = document.getElementById('mobile-menu-backdrop')

function toggleMenu(isOpen) {
    menu.classList.toggle('translate-x-full', !isOpen)
    backdrop.classList.toggle('opacity-100', isOpen)
    backdrop.classList.toggle('pointer-events-auto', isOpen)
    backdrop.classList.toggle('opacity-0', !isOpen)
    backdrop.classList.toggle('pointer-events-none', !isOpen)
    document.body.style.overflow = isOpen ? 'hidden' : ''
}

openBtn.addEventListener('click', () => toggleMenu(true))
closeBtn.addEventListener('click', () => toggleMenu(false))
backdrop.addEventListener('click', () => toggleMenu(false))
