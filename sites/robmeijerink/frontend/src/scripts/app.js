import '@fontsource-variable/inter';
import '../styles/app.css'

import { initTypedCursor } from './typed-cursor.js'
import { initReveal } from './reveal.js'
import { initSectionNav } from './section-nav.js'
import { initTimelineProgress } from './timeline-progress.js'
import { initTopbar } from './topbar.js'
import { initScrollWords } from './scroll-words.js'
import { initStackSections } from './stack-sections.js'
import { initPhotoMorph } from './photo-morph.js'
import { initCardSpotlight } from './card-spotlight.js'
import { initWhatsappReveal } from './whatsapp-reveal.js'

document.addEventListener('DOMContentLoaded', () => {
    initTypedCursor('Geen hypes, gewoon betrouwbare en stabiele software.')
    initReveal()
    initSectionNav()
    initTimelineProgress()
    initTopbar()
    initScrollWords()
    initStackSections()
    initCardSpotlight()
    initWhatsappReveal()
    if ('requestIdleCallback' in window) {
        requestIdleCallback(initPhotoMorph, { timeout: 2000 })
    } else {
        setTimeout(initPhotoMorph, 500)
    }
})
