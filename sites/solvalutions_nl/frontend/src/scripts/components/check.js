export const initCheck = () => {
    const root = document.querySelector('[data-check]');

    if (!root) return;

    const boxes = Array.from(root.querySelectorAll('input[type="checkbox"]'));
    const results = Array.from(root.querySelectorAll('[data-res]'));
    const out = root.querySelector('.check-out');
    const mail = root.querySelector('[data-mail]');
    const baseHref = mail.getAttribute('href');

    const update = () => {
        const picked = boxes.filter((box) => box.checked);
        const services = new Set();

        picked.forEach((box) => box.closest('[data-svc]').dataset.svc.split(' ').forEach((key) => services.add(key)));

        results.forEach((item) => {
            item.hidden = !services.has(item.dataset.res);
        });

        out.classList.toggle('has-res', picked.length > 0);
        mail.hidden = picked.length === 0;

        const lines = picked.map((box) => `- ${box.value}`).join('\n');
        const body = `Hoi Rob,\n\nIk herken dit:\n${lines}\n\nKunnen we kort bellen?`;

        mail.setAttribute('href', `${baseHref}&body=${encodeURIComponent(body)}`);
    };

    boxes.forEach((box) => box.addEventListener('change', update));
    update();
};
