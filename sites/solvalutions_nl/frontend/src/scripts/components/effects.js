const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches;

const initReveal = () => {
    document.querySelectorAll('[data-reveal-group]').forEach((group) => {
        Array.from(group.children).forEach((child, index) => {
            child.setAttribute('data-reveal', '');
            child.style.setProperty('--i', index);
        });
    });

    const targets = document.querySelectorAll('[data-reveal]');

    if (reduced || !('IntersectionObserver' in window)) {
        targets.forEach((target) => target.removeAttribute('data-reveal'));
        return;
    }

    const observer = new IntersectionObserver((entries) => {
        entries.forEach((entry) => {
            if (!entry.isIntersecting) return;

            entry.target.classList.add('is-in');
            observer.unobserve(entry.target);

            setTimeout(() => {
                entry.target.removeAttribute('data-reveal');
                entry.target.classList.remove('is-in');
                entry.target.style.removeProperty('--i');
            }, 1400);
        });
    }, { threshold: 0.12, rootMargin: '0px 0px -6% 0px' });

    targets.forEach((target) => observer.observe(target));
};

const initHeader = () => {
    const header = document.getElementById('main-nav');

    if (!header) return;

    const update = () => header.classList.toggle('is-scrolled', window.scrollY > 8);

    window.addEventListener('scroll', update, { passive: true });
    update();
};

const initProgress = () => {
    const bar = document.querySelector('.scroll-progress');

    if (!bar) return;

    const update = () => {
        const max = document.documentElement.scrollHeight - window.innerHeight;

        bar.style.setProperty('--sp', max > 0 ? (window.scrollY / max).toFixed(4) : 0);
    };

    window.addEventListener('scroll', update, { passive: true });
    window.addEventListener('resize', update);
    update();
};

const initWhatsapp = () => {
    const button = document.querySelector('a[href^="https://wa.me"]');

    if (!button || !window.matchMedia('(max-width: 47.99rem)').matches) return;

    let last = window.scrollY;

    button.style.transition = 'transform 0.35s ease, opacity 0.35s ease';

    window.addEventListener('scroll', () => {
        const current = window.scrollY;

        if (Math.abs(current - last) < 12) return;

        button.classList.toggle('wa-hide', current > last && current > 240);
        last = current;
    }, { passive: true });
};

export const initEffects = () => {
    initWhatsapp();
    initProgress();
    initReveal();
    initHeader();
};
