export const initCheck = () => {
    const root = document.querySelector('[data-check]');

    if (!root) return;

    const boxes = Array.from(root.querySelectorAll('input[type="checkbox"]'));
    const results = Array.from(root.querySelectorAll('[data-res]'));
    const out = root.querySelector('.check-out');
    const mail = root.querySelector('[data-mail]');
    const foot = root.querySelector('[data-foot]');
    const count = root.querySelector('[data-count]');
    const wa = root.querySelector('[data-wa]');
    const baseHref = mail.getAttribute('href');
    const waHref = wa.getAttribute('href');

    const update = () => {
        const picked = boxes.filter((box) => box.checked);
        const services = new Set();

        picked.forEach((box) => box.closest('[data-svc]').dataset.svc.split(' ').forEach((key) => services.add(key)));

        results.forEach((item) => {
            item.classList.toggle('is-on', services.has(item.dataset.res));
        });

        out.classList.toggle('has-res', picked.length > 0);
        foot.hidden = picked.length === 0;
        count.innerHTML = picked.length ? `<b>${picked.length}</b> van ${boxes.length} herkend` : `0 van ${boxes.length} herkend`;

        const lines = picked.map((box) => `- ${box.value}`).join('\n');
        const body = `Hoi Rob,\n\nIk kom via je website. Dit herken ik:\n${lines}\n\nKunnen we kort bellen?`;

        mail.setAttribute('href', `${baseHref}&body=${encodeURIComponent(body)}`);
        wa.setAttribute('href', `${waHref}?text=${encodeURIComponent(body)}`);
    };

    const jump = document.querySelector('[data-jump]');
    const jumpText = document.querySelector('[data-jump-text]');
    let outVisible = false;

    const syncJump = () => {
        if (!jump) return;

        const picked = boxes.filter((box) => box.checked).length;

        jump.hidden = picked === 0 || outVisible;
        jumpText.textContent = `${picked} herkend, bekijk wat ik kan doen`;
    };

    if (jump && 'IntersectionObserver' in window) {
        new IntersectionObserver(([entry]) => {
            outVisible = entry.isIntersecting;
            syncJump();
        }, { threshold: 0.25 }).observe(out);
    }

    boxes.forEach((box) => box.addEventListener('change', syncJump));
    boxes.forEach((box) => box.addEventListener('change', update));
    update();
};
