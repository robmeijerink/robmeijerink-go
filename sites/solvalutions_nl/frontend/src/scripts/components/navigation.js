export const initNavigation = () => {
    const nav = document.getElementById('main-nav');

    if (!nav) return;

    initDropdowns(nav);

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

const initDropdowns = (nav) => {
    const dropdowns = Array.from(nav.querySelectorAll('[data-dropdown]'));

    const close = (dropdown) => {
        dropdown.classList.remove('is-open');
        dropdown.querySelector('button').setAttribute('aria-expanded', 'false');
    };

    dropdowns.forEach((dropdown) => {
        const button = dropdown.querySelector('button');

        button.addEventListener('click', () => {
            const open = !dropdown.classList.contains('is-open');

            dropdowns.forEach(close);
            dropdown.classList.toggle('is-open', open);
            button.setAttribute('aria-expanded', String(open));
        });

        dropdown.addEventListener('focusout', (event) => {
            if (!dropdown.contains(event.relatedTarget)) close(dropdown);
        });
    });

    document.addEventListener('click', (event) => {
        dropdowns.forEach((dropdown) => {
            if (!dropdown.contains(event.target)) close(dropdown);
        });
    });

    document.addEventListener('keydown', (event) => {
        if (event.key !== 'Escape') return;

        dropdowns.forEach((dropdown) => {
            if (!dropdown.classList.contains('is-open')) return;

            close(dropdown);
            dropdown.querySelector('button').focus();
        });
    });
};
