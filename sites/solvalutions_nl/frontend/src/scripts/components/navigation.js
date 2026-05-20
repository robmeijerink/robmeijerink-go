export const initNavigation = () => {
    const nav = document.getElementById('main-nav');
    if (!nav) return;

    // --- 1. Original Scroll Behavior ---
    let lastScrollY = window.scrollY;
    let isNavVisible = true;
    let scrollUpAccumulator = 0;

    const handleScroll = () => {
        const currentScrollY = window.scrollY;
        const delta = currentScrollY - lastScrollY;

        if (currentScrollY > 50) {
            nav.classList.remove('bg-transparent', 'border-transparent', 'py-2');
            nav.classList.add('bg-canvas/10', 'backdrop-blur-sm', 'border-content-strong/5', 'py-0');
        } else {
            nav.classList.add('bg-transparent', 'border-transparent', 'py-2');
            nav.classList.remove('bg-canvas/10', 'backdrop-blur-sm', 'border-content-strong/5', 'py-0');
        }

        if (currentScrollY <= 150) {
            if (!isNavVisible) {
                nav.style.transform = 'translateY(0)';
                isNavVisible = true;
            }
            scrollUpAccumulator = 0;
        } else if (delta > 0) {
            if (isNavVisible) {
                nav.style.transform = 'translateY(-100%)';
                isNavVisible = false;
            }
            scrollUpAccumulator = 0;
        } else if (delta < 0) {
            scrollUpAccumulator += Math.abs(delta);
            if (!isNavVisible && scrollUpAccumulator > 80) {
                nav.style.transform = 'translateY(0)';
                isNavVisible = true;
            }
        }
        lastScrollY = currentScrollY;
    };

    window.addEventListener('scroll', handleScroll, { passive: true });

    // --- 2. New Mobile Drawer Toggle ---
    const btn = document.getElementById('mobile-menu-btn');
    const overlay = document.getElementById('mobile-overlay');
    const drawer = document.getElementById('mobile-drawer');

    if (btn && overlay && drawer) {
        const toggleMenu = () => {
            const isCurrentlyOpen = overlay.classList.contains('is-open');
            overlay.classList.toggle('is-open');
            drawer.classList.toggle('is-open');

            // A11y update
            btn.setAttribute('aria-expanded', !isCurrentlyOpen);
        };

        btn.addEventListener('click', toggleMenu);
        overlay.addEventListener('click', toggleMenu);
    }
};
