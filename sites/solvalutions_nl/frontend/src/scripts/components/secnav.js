export const initSecnav = () => {
    const nav = document.querySelector('.secnav');

    if (!nav || !('IntersectionObserver' in window)) return;

    const links = Array.from(nav.querySelectorAll('[data-sec]'));
    const sections = links.map((link) => document.getElementById(link.dataset.sec)).filter(Boolean);
    const dark = new Set(['methodology-section', 'cta-section']);

    const observer = new IntersectionObserver((entries) => {
        entries.forEach((entry) => {
            if (!entry.isIntersecting) return;

            links.forEach((link) => {
                const current = link.dataset.sec === entry.target.id;

                link.classList.toggle('is-current', current);

                if (current) link.setAttribute('aria-current', 'true');
                else link.removeAttribute('aria-current');
            });

            nav.classList.toggle('on-dark', dark.has(entry.target.id));
        });
    }, { rootMargin: '-45% 0px -50% 0px' });

    sections.forEach((section) => observer.observe(section));
};
