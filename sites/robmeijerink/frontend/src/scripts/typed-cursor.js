export function initTypedCursor() {
    const text = 'Geen hypes, gewoon betrouwbare en stabiele software.'
    const target = document.getElementById('typed-manifesto')
    const cursor = document.getElementById('cursor')

    if (!target) return

    let i = 0

    function typeWriter() {
        if (i < text.length) {
            target.textContent += text.charAt(i)
            i++
            setTimeout(typeWriter, 40)
        } else {
            setTimeout(() => {
                if (cursor) cursor.classList.add('fade-out')
            }, 3000)
        }
    }

    // Start de animatie
    setTimeout(typeWriter, 1250)
}
