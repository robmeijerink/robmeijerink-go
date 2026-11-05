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
    let scrollHidden = false;
    let idleTimer;
    const zonesInView = new Set();

    const render = () => {
        button.classList.toggle('wa-hide', scrollHidden || zonesInView.size > 0);
    };

    const wake = () => {
        button.classList.remove('wa-idle');
        clearTimeout(idleTimer);
        idleTimer = setTimeout(() => button.classList.add('wa-idle'), 1800);
    };

    button.style.transition = 'transform 0.35s ease, opacity 0.35s ease, filter 0.35s ease';

    // The call to action and the footer already show phone and mail, so the button steps aside there.
    if ('IntersectionObserver' in window) {
        const observer = new IntersectionObserver((entries) => {
            entries.forEach((entry) => {
                if (entry.isIntersecting) zonesInView.add(entry.target);
                else zonesInView.delete(entry.target);
            });
            render();
        });

        document.querySelectorAll('#cta-section, #main-footer').forEach((zone) => observer.observe(zone));
    }

    window.addEventListener('scroll', () => {
        const current = window.scrollY;

        wake();

        if (Math.abs(current - last) < 12) return;

        scrollHidden = current > last && current > 240;
        render();
        last = current;
    }, { passive: true });

    window.addEventListener('touchstart', wake, { passive: true });
    wake();
};

const initFaqMore = () => {
    document.querySelectorAll('.faq').forEach((faq) => {
        const items = faq.querySelectorAll('details');

        if (items.length <= 6) return;

        faq.classList.add('faq-collapsed');

        const wrap = document.createElement('p');
        wrap.className = 'faq-more';

        const button = document.createElement('button');
        button.type = 'button';
        button.className = 'btn btn-secondary focus-ring';
        button.setAttribute('aria-expanded', 'false');
        button.textContent = `Toon alle ${items.length} vragen`;

        button.addEventListener('click', () => {
            faq.classList.remove('faq-collapsed');
            wrap.remove();
        });

        wrap.appendChild(button);
        faq.after(wrap);
    });
};

export const initEffects = () => {
    initWhatsapp();
    initFaqMore();
    initProgress();
    initReveal();
    initHeader();
};
