export const initNavigation = () => {
    const nav = document.getElementById('main-nav');

    if (!nav) return;

    const btn = document.getElementById('mobile-menu-btn');
    const overlay = document.getElementById('mobile-overlay');
    const drawer = document.getElementById('mobile-drawer');

    if (!btn || !overlay || !drawer) return;

    const toggleMenu = () => {
        const isCurrentlyOpen = overlay.classList.contains('is-open');

        nav.classList.toggle('is-open');
        btn.classList.toggle('is-open');
        overlay.classList.toggle('is-open');
        drawer.classList.toggle('is-open');

        btn.setAttribute('aria-expanded', String(!isCurrentlyOpen));

        document.body.style.overflow = isCurrentlyOpen ? '' : 'hidden';
    };

    btn.addEventListener('click', toggleMenu);
    overlay.addEventListener('click', toggleMenu);

    document.addEventListener('keydown', (event) => {
        if (event.key === 'Escape' && overlay.classList.contains('is-open')) toggleMenu();
    });

    nav.addEventListener('click', (event) => {
        if (!nav.classList.contains('is-open') || event.target.closest('button, a')) return;

        if (event.clientX < drawer.getBoundingClientRect().left) toggleMenu();
    });
};
