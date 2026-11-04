export function initTopbar() {
    const topbar = document.querySelector('[data-topbar]')

    if (!topbar) return

    function update() {
        topbar.classList.toggle('is-scrolled', window.scrollY > 8)
    }

    window.addEventListener('scroll', update, { passive: true })
    update()
}
