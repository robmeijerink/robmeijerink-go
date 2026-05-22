export const initNavigation = () => {
    const nav = document.getElementById('main-nav');
    const navBackground = document.getElementById('nav-background');
    if (!nav || !navBackground) return;

    let lastScrollY = window.scrollY;
    let isNavVisible = true;
    let scrollUpAccumulator = 0;

    const handleScroll = () => {
        const currentScrollY = window.scrollY;
        const delta = currentScrollY - lastScrollY;

        if (currentScrollY > 50) {
            navBackground.classList.remove('bg-transparent', 'border-transparent');
            navBackground.classList.add('bg-canvas/10', 'backdrop-blur-sm', 'border-content-strong/5');
        } else {
            navBackground.classList.remove('bg-canvas/10', 'backdrop-blur-sm', 'border-content-strong/5');
            navBackground.classList.add('bg-transparent', 'border-transparent');
        }

        if (currentScrollY <= 150) {
            if (!isNavVisible) {
                nav.style.transform = '';
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
                nav.style.transform = '';
                isNavVisible = true;
            }
        }
        lastScrollY = currentScrollY;
    };

    window.addEventListener('scroll', handleScroll, { passive: true });

    const btn = document.getElementById('mobile-menu-btn');
    const overlay = document.getElementById('mobile-overlay');
    const drawer = document.getElementById('mobile-drawer');

    if (btn && overlay && drawer) {
        const toggleMenu = () => {
            const isCurrentlyOpen = overlay.classList.contains('is-open');

            nav.classList.toggle('is-open');
            btn.classList.toggle('is-open');
            overlay.classList.toggle('is-open');
            drawer.classList.toggle('is-open');

            btn.setAttribute('aria-expanded', !isCurrentlyOpen);

            document.body.style.overflow = isCurrentlyOpen ? '' : 'hidden';
        };

        btn.addEventListener('click', toggleMenu);
        overlay.addEventListener('click', toggleMenu);

        nav.addEventListener('click', (e) => {
            if (nav.classList.contains('is-open') && !e.target.closest('button, a')) {
                const drawerLeft = drawer.getBoundingClientRect().left;
                if (e.clientX < drawerLeft) {
                    toggleMenu();
                }
            }
        });
    }
};
