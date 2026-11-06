export const initSceneFocus = () => {
    const stage = document.querySelector('[data-scene]');

    if (!stage) return;

    const pins = Array.from(stage.querySelectorAll('.sc-dot'));
    const first = pins[0];
    const desktop = window.matchMedia('(min-width: 64rem)');
    let touched = false;

    const set = (value) => {
        if (value) stage.setAttribute('data-focus', value);
        else stage.removeAttribute('data-focus');
    };

    const close = () => pins.forEach((pin) => pin.classList.remove('is-open'));

    const settle = () => {
        if (touched || !desktop.matches || !first) return;

        first.classList.add('is-open');
        set(first.getAttribute('data-focus'));
    };

    pins.forEach((pin) => {
        const value = pin.getAttribute('data-focus');

        pin.addEventListener('pointerenter', () => {
            touched = true;
            close();
            set(value);
        });
        pin.addEventListener('pointerleave', () => set(null));
        pin.addEventListener('focus', () => {
            touched = true;
            close();
            set(value);
        });
        pin.addEventListener('blur', () => set(null));
        pin.addEventListener('click', () => {
            touched = true;
            close();
            pin.classList.add('is-open');
            set(value);
        });
    });

    document.addEventListener('click', (event) => {
        if (!event.target.closest('.sc-dot')) {
            close();
            set(null);
        }
    });

    document.querySelectorAll('#hero-section .pil').forEach((item) => {
        const value = item.getAttribute('data-focus');

        item.addEventListener('pointerenter', () => set(value));
        item.addEventListener('pointerleave', () => set(null));
    });

    if (!window.matchMedia('(prefers-reduced-motion: reduce)').matches) setTimeout(settle, 1200);
    else settle();
};
